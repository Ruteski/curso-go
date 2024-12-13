﻿package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket *s3.Bucket
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(
				"ASDFASDFASDFASFD",
				"ASDFASdfasregq35ghq3%E",
				"",
			),
		},
	)
	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-exemplo"
}

func main() {

}