package bucket

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Inputs struct {
	Key        string
	Secret     string
	Region     string
	BucketName string
}

func (in Inputs) Run() {
	if in.Key == "" || in.Secret == "" {
		fmt.Println("Verify your credentials saved! Not received.")
		return
	}
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(in.Region),
		Credentials: credentials.NewStaticCredentials(in.Key, in.Secret, ""),
	})
	if err != nil {
		fmt.Println("Failed to create session, verify credentials")
		return
	}
	svc := s3.New(sess)
	in.runCreate(svc)
}

func (in Inputs) runCreate(svc *s3.S3) {
	if in.BucketName == "" {
		fmt.Println("Not received bucket name!")
		return
	}
	bn := &s3.CreateBucketInput{
		Bucket: aws.String(in.BucketName),
	}
	_, err := svc.CreateBucket(bn)
	if err != nil {
		fmt.Printf("Failed create bucket. error: %v\n", err)
		return
	}
	fmt.Println("Bucket created")
}