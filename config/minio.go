package config

import (
	"context"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinioClient() *minio.Client {
	logger.Debugf("connecting to MinIO server at %s", Envs.MINIO_ENDPOINT)

	minioClient, err := minio.New(Envs.MINIO_ENDPOINT, &minio.Options{
		Creds:  credentials.NewStaticV4(Envs.MINIO_ACCESS_KEY, Envs.MINIO_SECRET_KEY, ""),
		Secure: false,
	})
	if err != nil {
		logger.Fatalf("failed to initialize MinIO client: %v", err)
	}

	// verify connection
	err = minioClient.MakeBucket(context.Background(), "test", minio.MakeBucketOptions{})
	if err != nil {
		logger.Fatalf("failed to list buckets: %v", err)
	}
	// logger.Infof("successfully connected to MinIO. found %d buckets.", len(buckets))

	return minioClient
}
