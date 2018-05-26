package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"golang.org/x/sync/errgroup"
)

const (
	testdataPath = "../testdata"
	region       = "ap-northeast-1"
	s3Bucket     = "sdk-go-test"
	timeoutSec   = 10
)

/*
dd if=/dev/zero of=1m bs=1M count=1
dd if=/dev/zero of=10M bs=1M count=10
dd if=/dev/zero of=100M bs=1M count=100
dd if=/dev/zero of=200M bs=1M count=200
*/
var files = []string{"1M", "10M", "100M", "200M"}

func putS3ObjectWithContext(ctx context.Context, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return err
	}
	svc := s3.New(sess)
	params := &s3.PutObjectInput{
		Body:   f,
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(file),
	}

	if _, err := svc.PutObjectWithContext(ctx, params, func(r *request.Request) {
		start := time.Now()
		r.Handlers.Complete.PushBack(func(req *request.Request) {
			fmt.Printf("request %s took %s to complete\n", req.RequestID, time.Since(start))
		})
	}); err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
	}

	fmt.Printf("successfully uploaded file to %s/%s\n", s3Bucket, file)
	return nil
}

func main() {
	// Concurrent execution
	eg := errgroup.Group{}

	for _, file := range files {
		filePath := filepath.Join(testdataPath, file)
		fmt.Printf("upload %s...\n", filePath)
		ctx, cancel := context.WithTimeout(context.Background(), timeoutSec*time.Second)
		defer cancel()

		eg.Go(func() error {
			if err := putS3ObjectWithContext(ctx, filePath); err != nil {
				return err
			}

			return nil
		})

	}

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
