package bucket

import (
	"aws/bucket/pkg/prompt"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jedib0t/go-pretty/table"
	"os"
)

type Inputs struct {
	Key        string
	Secret     string
	Region     string
	BucketName string
	Command    string
}

const (
	list   = "list"
	delete = "delete"
	create = "create"
)

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
	switch in.Command {
	case list:
		in.runList(svc)
	case create:
		in.runCreate(svc)
	case delete:
		in.runDelete(svc)
	default:
		fmt.Printf("Command (%s) not found\n", in.Command)
	}
}

func (in Inputs) runList(svc *s3.S3) {
	res, err := in.list(svc)
	if err != nil {
		fmt.Printf("Error list bucket, error: %v", err)
	} else {
		printList(res)
	}
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

func printList(r *s3.ListBucketsOutput) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Count", "Bucket Name", "Creation"})
	for i, b := range r.Buckets {
		t.AppendRows([]table.Row{
			{i+1, aws.StringValue(b.Name), aws.TimeValue(b.CreationDate)},
		})
	}
	t.SetStyle(table.StyleLight)
	t.Render()
}