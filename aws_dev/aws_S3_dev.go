package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"os"
	"time"
    "strings"

)

// https://www.socketloop.com/tutorials/golang-setting-up-configure-aws-credentials-with-official-aws-sdk-go
// https://github.com/aws/aws-sdk-go/wiki/common-examples

// A Object provides details of an S3 object
type Object struct {
	Bucket         string
	Key            string
	Encrypted      bool
	EncryptionType string
}

// An ErrObject provides details of the error occurred retrieving
// an object's status.
type ErrObject struct {
	Bucket string
	Key    string
	Error  error
}

// A Bucket provides details about a bucket and its objects
type Bucket struct {
	Owner        string
	Name         string
	CreationDate time.Time
	Region       string
	Objects      []Object
	Error        error
	ErrObjects   []ErrObject
}




func main() {

	// load env vars
	creds := credentials.NewEnvCredentials()

	credValue, err := creds.Get()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Value of credentials : ", credValue)
	fmt.Println("----------------------------------------------")

	fmt.Println("Raw credentials : ", creds)

	sess := session.New(&aws.Config{Region: aws.String("eu-west-1")})
	s3client := s3.New(sess)

	bucketName := "eddysstuff"
	key := "TestFile.txt"

	result, err := s3client.CreateBucket(&s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	if err != nil {
		fmt.Println("Failed to create bucket", err)
		return
	}

	fmt.Println("CreateBucket result : ", result.GoString())

	uploadResult, err2 := s3client.PutObject(&s3.PutObjectInput{
		Body:   strings.NewReader("Hello World!"),
		Bucket: &bucketName,
		Key:    &key,
	})

	if err2 != nil {
		fmt.Printf("Failed to upload data to %s/%s, %s\n", bucketName, key, err)
		return
	}

	fmt.Println("CreateBucket result : ", uploadResult.GoString())


	// https://github.com/aws/aws-sdk-go/blob/master/service/s3/examples_test.go#L860-L879

	var params2 *s3.ListBucketsInput

	resp2, err := s3client.ListBuckets(params2)

	fmt.Println(resp2)

}

