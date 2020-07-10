package bucket

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/jedib0t/go-pretty/table"
)

type Inputs struct {
	Key    string
	Secret string
	Region string
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
	in.runList(svc)
}

func (in Inputs) runList(svc *s3.S3) {
	res, err := in.list(svc)
	if err != nil {
		fmt.Printf("Error list bucket, error: %v", err)
	} else {
		printList(res)
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
			{i + 1, aws.StringValue(b.Name), aws.TimeValue(b.CreationDate)},
		})
	}
	t.SetStyle(table.StyleLight)
	t.Render()
}
