package domain

import r "api/utils/result"

type FileUploader interface {
	UploadCoworkingImage(coworking *Coworking, imageName string, imageData []byte) r.Result[string]
}
