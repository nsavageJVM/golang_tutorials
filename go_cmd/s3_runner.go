package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"io/ioutil"
//	"os"

)


var (

	s3filePath    = kingpin.Flag("s3file", "path to the file").String()
	uploadFile      = kingpin.Command("upload", "Upload a file to S3")
	downloadFile      = kingpin.Command("download", "Download a file to S3")
	listS3      = kingpin.Command("list", "List S3 buckets")
	createS3      = kingpin.Command("create", "Create S3 buckets")

)

func main() {

	kingpin.Version("0.0.1")
	kingpin.Parse()

	sess := session.New(&aws.Config{Region: aws.String("eu-west-1")})
	s3client := s3.New(sess)


	switch kingpin.Parse() {
	// Upload
	case uploadFile.FullCommand():
		fmt.Printf("uploading file: %s to S3 \n ", *s3filePath)
		dat, err := ioutil.ReadFile(*s3filePath)
		check(err)
		fmt.Print(string(dat))


	// Download
	case downloadFile.FullCommand():
		fmt.Printf("downloading file: %s from S3 \n ", *s3filePath)

	// create S3
	case createS3.FullCommand():

		fmt.Printf("create S3 bucket: %s from S3 \n ", *s3filePath)

	// list
	case listS3.FullCommand():
		var params2 *s3.ListBucketsInput
		resp2, err := s3client.ListBuckets(params2)
		check(err)
		fmt.Println(resp2)

	}


}

func check(err error) {
	if err != nil {
		fmt.Println("S3 Failed  ", err)
		return
	}

}

