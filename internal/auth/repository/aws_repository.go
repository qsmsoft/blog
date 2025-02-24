package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/qsmsoft/blog/internal/auth"
	"github.com/qsmsoft/blog/internal/models"
)

// authAWSRepository is AWS S3 repository
type authAWSRepository struct {
	client *minio.Client
}

// NewAuthAWSRepository is AWS S3 repository constructor
func NewAuthAWSRepository(awsClient *minio.Client) auth.AWSRepository {
	return &authAWSRepository{client: awsClient}
}

// PutObject uploads file to AWS
func (aws *authAWSRepository) PutObject(ctx context.Context, input models.UploadInput) (*minio.UploadInfo, error) {
	options := minio.PutObjectOptions{
		ContentType:  input.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	uploadInfo, err := aws.client.PutObject(ctx, input.BucketName, aws.generateFileName(input.Name), input.File, input.Size, options)
	if err != nil {
		return nil, errors.Wrap(err, "authAWSRepository.FileUpload.PutObject")
	}

	return &uploadInfo, err
}

// GetObject downloads file from AWS
func (aws *authAWSRepository) GetObject(ctx context.Context, bucket string, fileName string) (*minio.Object, error) {
	object, err := aws.client.GetObject(ctx, bucket, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "authAWSRepository.FileDownload.GetObject")
	}
	return object, nil
}

// RemoveObject deletes file from AWS
func (aws *authAWSRepository) RemoveObject(ctx context.Context, bucket string, fileName string) error {
	if err := aws.client.RemoveObject(ctx, bucket, fileName, minio.RemoveObjectOptions{}); err != nil {
		return errors.Wrap(err, "authAWSRepository.RemoveObject")
	}
	return nil
}

// generateFileName generates file name
func (aws *authAWSRepository) generateFileName(fileName string) string {
	uid := uuid.New().String()
	return fmt.Sprintf("%s-%s", uid, fileName)
}
