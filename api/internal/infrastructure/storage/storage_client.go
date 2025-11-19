package storage

import (
	"context"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3StorageClient struct {
	s3Client *s3.Client
	bucket   string
	endpoint string
}

func NewStorageClient(ctx context.Context, accountID, accessKey, secretKey, bucket string) (*S3StorageClient, error) {
	endpoint := "https://" + accountID + ".r2.cloudflarestorage.com"

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("auto"),
		config.WithCredentialsProvider(
			credentials.NewStaticCredentialsProvider(accessKey, secretKey, ""),
		),
	)

	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)
	})

	return &S3StorageClient{
		s3Client: client,
		bucket:   bucket,
		endpoint: endpoint,
	}, nil
}

func (s *S3StorageClient) UploadFile(filename string, fileHeader *multipart.FileHeader, ctx context.Context) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	_, err = s.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		return "", err
	}

	return s.endpoint + "/" + filename, nil
}

func (s *S3StorageClient) Ping(ctx context.Context) error {
	_, err := s.s3Client.HeadBucket(ctx, &s3.HeadBucketInput{
		Bucket: aws.String(s.bucket),
	})
	return err
}
