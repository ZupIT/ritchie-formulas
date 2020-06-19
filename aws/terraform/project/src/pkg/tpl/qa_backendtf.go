package tpl

const (
	QABackendtf = `
region = "{{.BucketRegion}}"
bucket = "{{.BucketName}}"
key    = "terraform.tfstate"
`
)
