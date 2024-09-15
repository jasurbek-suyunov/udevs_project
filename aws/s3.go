package aws

import (
	"fmt"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	bucket = "jasurdev" // S3'dagi bucket nomingiz
)

func UploadFileToS3(file multipart.File, fileName string) error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"), // O'zingizning mintaqangiz
	})
	if err != nil {
		return fmt.Errorf("failed to create session: %v", err)
	}

	uploader := s3.New(sess)

	// Faylni S3'ga yuklash
	_, err = uploader.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("test/" + fileName), // Faylni 'test' papkasiga yuklash
		Body:   file,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %v", err)
	}

	return nil
}		