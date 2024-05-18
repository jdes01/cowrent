package aws

import (
	domain "api/coworkings/domain"
	r "api/utils/result"
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type FileUploader struct {
	BUCKET_NAME string
	s3          *s3.S3
}

func NewFileUploader(s3 *s3.S3) *FileUploader {
	return &FileUploader{
		s3:          s3,
		BUCKET_NAME: "coworking-images",
	}
}

func (uploader *FileUploader) UploadCoworkingImage(coworking *domain.Coworking, imageName string, imageData []byte) r.Result[string] {

	bucketPath := fmt.Sprintf("%s-images/", coworking.UUID.Value) + imageName

	_, err := uploader.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(uploader.BUCKET_NAME),
		Key:    aws.String(bucketPath),
		Body:   bytes.NewReader(imageData),
	})

	if err != nil {
		return r.NewResult("", err)
	}

	imagePath := fmt.Sprintf("%s/%s", uploader.BUCKET_NAME, bucketPath)

	return r.NewResult(imagePath, nil)
}
