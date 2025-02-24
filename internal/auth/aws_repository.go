package auth

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/qsmsoft/blog/internal/models"
)

type AWSRepository interface {
	PutObject(ctx context.Context, input models.UploadInput) (*minio.UploadInfo, error)
	GetObject(ctx context.Context, bucket string, fileName string) (*minio.Object, error)
	RemoveObject(ctx context.Context, bucket string, fileName string) error
}
