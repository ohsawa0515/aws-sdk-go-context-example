package main

import (
	"context"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	region     = "ap-northeast-1"
	s3Bucket   = "sdk-go-test2"
	timeoutSec = 120
)

func createBucketWithContext(ctx context.Context) error {
	sess, err := session.NewSession(&aws.Config{
		Region:     aws.String(region),
		MaxRetries: aws.Int(3),
	})
	if err != nil {
		return err
	}

	svc := s3.New(sess)
	if _, err := svc.CreateBucketWithContext(ctx, &s3.CreateBucketInput{
		Bucket: aws.String(s3Bucket),
	}); err != nil {
		return err
	}

	if err := svc.WaitUntilBucketExistsWithContext(ctx,
		&s3.HeadBucketInput{
			Bucket: aws.String(s3Bucket),
		},
		request.WithWaiterDelay(request.ConstantWaiterDelay(timeoutSec*time.Second)),
	); err != nil {
		return err
	}

	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := createBucketWithContext(ctx); err != nil {
		log.Fatal(err)
	}
}
