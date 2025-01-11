package repository

import (
	"backend/utils/helper"
	"context"
	"fmt"
	"mime/multipart"

	"github.com/minio/minio-go/v7"
)

type FileRepo struct {
	client *minio.Client
}

type IFileRepo interface {
	Upload(ctx context.Context, file *multipart.FileHeader, filename string, bucketname string) (string, error)
	GetUrl(ctx context.Context, filename string, bucketname string, download bool) (string, error)
	Delete(ctx context.Context, filename string, bucketname string) error
}

func NewFileRepo(client *minio.Client) IFileRepo {
	return &FileRepo{
		client: client,
	}
}

func (r *FileRepo) Upload(ctx context.Context, file *multipart.FileHeader, filename string, bucketname string) (string, error) {
	fileContent, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer fileContent.Close()

	_, err = r.client.PutObject(ctx, bucketname, filename, fileContent, file.Size, minio.PutObjectOptions{
		ContentType: helper.GetMimeType(filename),
	})
	if err != nil {
		return "", fmt.Errorf("failed to upload : %w", err)
	}

	return filename, nil
}

func (r *FileRepo) GetUrl(ctx context.Context, filename string, bucketname string, download bool) (string, error) {
	presignedUrl, err := r.client.PresignedGetObject(ctx, bucketname, filename, 60*60, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get url: %w", err)
	}
	return presignedUrl.String(), nil
}

func (r *FileRepo) Delete(ctx context.Context, filename string, bucketname string) error {
	err := r.client.RemoveObject(ctx, bucketname, filename, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}
