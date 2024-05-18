package aws_config

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/eldimious/golang-api-showcase/utils/env"

	aws_sdk "github.com/aws/aws-sdk-go/aws"
)

type Config struct {
	AWS_region   string
	AWS_endpoint string
	S3_service   *s3.S3
}

func NewConfig() *Config {
	env.CheckDotEnv()

	aws_region := env.MustGet("AWS_DEFAULT_REGION")
	aws_s3_endpoint_url := env.MustGet("AWS_S3_ENDPOINT_URL")

	FORCE_PATH_STYLE := true

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws_sdk.Config{
			Region:           aws_sdk.String(aws_region),
			Endpoint:         aws_sdk.String(aws_s3_endpoint_url),
			S3ForcePathStyle: &FORCE_PATH_STYLE,
		},
	}))

	s3_service := s3.New(sess)
	return &Config{S3_service: s3_service}
}
