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
	in.runDelete(svc)
}

func (in Inputs) runDelete(svc *s3.S3) {
	res, err := in.list(svc)
	if err != nil {
		fmt.Printf("Error list bucket to delete, error: %v", err)
	}
	var bItems []string
	for _, b := range res.Buckets {
		bItems = append(bItems, aws.StringValue(b.Name))
	}
	if len(bItems) == 0 {
		fmt.Printf("Not found bucket to delete")
		return
	}
	bSelect, _ := prompt.List("Select bucket to delete: ", bItems)
	cItems := []string{"NO", "YES"}
	c, _ := prompt.List(fmt.Sprintf("Confirm delete bucket name: %s", bSelect), cItems)
	switch c {
	case "NO":
		fmt.Printf("Bucket %s not deleted\n", bSelect)
	case "YES":
		bn := &s3.DeleteBucketInput{
			Bucket: aws.String(bSelect),
		}
		_, err := svc.DeleteBucket(bn)
		if err != nil {
			fmt.Printf("Error delete bucket %s, error: %v\n", bSelect, err)
			return
		}
		fmt.Printf("Bucket %s deleted.\n", bSelect)
	}
}

func (in Inputs) list(svc *s3.S3) (*s3.ListBucketsOutput, error) {
	result, err := svc.ListBuckets(nil)
	if err != nil {
		return nil, fmt.Errorf("Failed list bucket, error: %s\n", err.Error())
	}
	return result, nil
}
