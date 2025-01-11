package file_storage_util

import (
	"backend/utils/helper"
	"context"
	"fmt"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/minio/minio-go/v7"
)

type FileStorageUtil struct {
	client *minio.Client
}

type IFileStorageUtil interface {
	Upload(ctx context.Context, file *multipart.FileHeader, filename string, bucketname string) (string, error)
	GetUrl(ctx context.Context, filename string, bucketname string, download bool) (string, error)
	Delete(ctx context.Context, filename string, bucketname string) error
}

func NewFileStorageUtil(client *minio.Client) IFileStorageUtil {
	return &FileStorageUtil{
		client: client,
	}
}

func (r *FileStorageUtil) Upload(ctx context.Context, file *multipart.FileHeader, filename string, bucketname string) (string, error) {
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

func (r *FileStorageUtil) GetUrl(ctx context.Context, filename string, bucketname string, download bool) (string, error) {
	reqParams := url.Values{
		"response-content-type": []string{helper.GetMimeType(filename)},
	}
	if download {
		reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	} else {
		reqParams.Set("response-content-disposition", fmt.Sprintf("inline; filename=\"%s\"", filename))
	}

	presignedUrl, err := r.client.PresignedGetObject(ctx, bucketname, filename, time.Minute, reqParams)
	if err != nil {
		return "", fmt.Errorf("failed to get url: %w", err)
	}
	return presignedUrl.String(), nil
}

func (r *FileStorageUtil) Delete(ctx context.Context, filename string, bucketname string) error {
	err := r.client.RemoveObject(ctx, bucketname, filename, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete file: %w", err)
	}
	return nil
}
