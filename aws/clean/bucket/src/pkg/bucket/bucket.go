package bucket

import (
	"aws/bucket/pkg/prompt"
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
	Command    string
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
	in.runClean(svc)
}

func (in Inputs) runClean(svc *s3.S3) {
	res, err := in.list(svc)
	if err != nil {
		fmt.Printf("Error list bucket to clean, error: %v", err)
	}
	var bItems []string
	for _, b := range res.Buckets {
		bItems = append(bItems, aws.StringValue(b.Name))
	}
	if len(bItems) == 0 {
		fmt.Printf("Not found bucket to clean")
		return
	}
	bSelect, _ := prompt.List("Select bucket to clean: ", bItems)
	confirm, _ := prompt.List(fmt.Sprintf("Confirm clean bucket name: %s", bSelect), []string{"NO", "YES"})
	switch confirm {
	case "NO":
		fmt.Printf("Bucket %s not cleaned\n", bSelect)
	case "YES":
		fmt.Printf("Cleaning...")
		listObjectsInput := &s3.ListObjectsInput{
			Bucket:  aws.String(bSelect),
			MaxKeys: aws.Int64(20),
		}

		svc.ListObjectsPages(listObjectsInput, func(page *s3.ListObjectsOutput, lastPage bool) bool {
			for _, value := range page.Contents {
				deleteObjectInput := &s3.DeleteObjectInput{
					Bucket: aws.String(bSelect),
					Key:    value.Key,
				}

				_, err := svc.DeleteObject(deleteObjectInput)
				if err != nil {
					fmt.Printf("Error on delete object %s, error: %v\n", err)
					return false
				}

				fmt.Println("Deleted object: ", *value.Key)
			}
			return true
		})

		fmt.Printf("Bucket %s cleaned.\n", bSelect)
	}
}

func (in Inputs) list(svc *s3.S3) (*s3.ListBucketsOutput, error) {
	result, err := svc.ListBuckets(nil)
	if err != nil {
		return nil, fmt.Errorf("Failed list bucket, error: %s\n", err.Error())
	}
	return result, nil
}
