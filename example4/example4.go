package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	DefaultS3Bucket = "test-bucket"
	DefaultS3Key    = "path/"
	DefaultRegion   = "ap-northeast-1"
	DefaultMaxKeys  = 100
)

var region, s3Bucket, s3Key string
var maxKeys int64

func listObjectsPagesWithContext(ctx context.Context) error {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(region)})
	if err != nil {
		return err
	}
	svc := s3.New(sess)

	// Using context
	if err := svc.ListObjectsPagesWithContext(ctx,
		&s3.ListObjectsInput{
			Bucket:  aws.String(s3Bucket),
			Prefix:  aws.String(s3Key),
			MaxKeys: aws.Int64(maxKeys),
		},
		func(page *s3.ListObjectsOutput, lastPage bool) bool {
			fmt.Println("Received", len(page.Contents), "objects in page")
			for _, obj := range page.Contents {
				fmt.Println("Key:", aws.StringValue(obj.Key))
			}
			return true
		},
	); err != nil {
		return err
	}

	// Without using context
	if err := svc.ListObjectsPages(&s3.ListObjectsInput{
		Bucket:  aws.String(s3Bucket),
		Prefix:  aws.String(s3Key),
		MaxKeys: aws.Int64(maxKeys),
	},
		func(page *s3.ListObjectsOutput, lastPage bool) bool {
			fmt.Println("Received", len(page.Contents), "objects in page")
			for _, obj := range page.Contents {
				fmt.Println("Key:", aws.StringValue(obj.Key))
			}
			return true
		}); err != nil {
		return err
	}

	return nil
}

func main() {
	flag.StringVar(&s3Bucket, "b", DefaultS3Bucket, "Bucket name")
	flag.StringVar(&s3Key, "k", DefaultS3Key, "Key of s3 bucket")
	flag.StringVar(&region, "r", DefaultRegion, "Region")
	flag.Int64Var(&maxKeys, "m", DefaultMaxKeys, "Maximum number of keys to fetch")
	flag.Parse()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := listObjectsPagesWithContext(ctx); err != nil {
		fmt.Println(err)
	}
}
