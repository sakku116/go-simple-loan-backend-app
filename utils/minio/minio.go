package minio_util

import (
	"backend/domain/model"
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/op/go-logging"
)

var logger = logging.MustGetLogger("minio_util")

func EnsureBucketsFromModels(minioClient *minio.Client, models ...model.IModel) error {
	for _, model := range models {
		props := model.GetProps()
		bucketName := props.BucketName

		exists, err := minioClient.BucketExists(context.Background(), bucketName)
		if err != nil {
			logger.Errorf("failed to check if bucket %s exists: %v", bucketName, err)
			return err
		}

		if !exists {
			err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{})
			if err != nil {
				logger.Errorf("failed to create bucket %s: %v", bucketName, err)
				return err
			}
		}

		logger.Infof("bucket %s created", bucketName)
	}

	return nil
}
