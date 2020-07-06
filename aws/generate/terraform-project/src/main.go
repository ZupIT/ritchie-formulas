package main

import (
	"project/pkg/aws"
	"os"
)

func main() {
	name := os.Getenv("PROJECT_NAME")
	loc := os.Getenv("PROJECT_LOCATION")
	bucketName := os.Getenv("BUCKET_NAME")
	bucketRegion := os.Getenv("BUCKET_REGION")
	pwd := os.Getenv("CURRENT_PWD")

	aws.Input{
		ProjectName:     name,
		ProjectLocation: loc,
		BucketName: bucketName,
		BucketRegion: bucketRegion,
		PWD: pwd,
	}.Run()
}
