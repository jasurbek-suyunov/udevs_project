package handler

import (
	"fmt"
	"github.com/jasurbek-suyunov/udevs_project/config"
	"github.com/jasurbek-suyunov/udevs_project/models"
	"log"
	"mime/multipart"
	"net/url"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

var (
	s3Session *s3.S3
	bucket    string
)
// connect to S3
func connectS3(cnf *config.Config) error {
	var awss3_config = models.Amazons3Config{
		AccessKey: cnf.Amazons3AccessKey,
		SecretKey: cnf.Amazons3SecretKey,
		Region:    cnf.Amazons3Region,
		Bucket:    cnf.Amazons3Bucket,
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(awss3_config.Region),
		Credentials: credentials.NewStaticCredentials(awss3_config.AccessKey, awss3_config.SecretKey, ""),
	})
	if err != nil {
		log.Fatalf("Failed to create AWS session: %v", err)
		return err
	}

	s3Session = s3.New(sess)
	bucket = awss3_config.Bucket
	return nil
}

// upload file to S3
func uploadToS3(file multipart.File, fileHeader *multipart.FileHeader, folder string) (string, error) {
	key := fmt.Sprintf("%s/%s%s", folder, uuid.New().String(), filepath.Ext(fileHeader.Filename))
	_, err := s3Session.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
		ACL:    aws.String("private"),
	})
	if err != nil {
		return "", err
	}

	// return file url
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket, url.PathEscape(key)), nil
}