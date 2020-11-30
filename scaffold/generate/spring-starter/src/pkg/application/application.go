package application

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/ZupIT/ritchie-formulas/scaffold/spring-starter/src/pkg/file/fileutil"
)

const (
	starterURL = "https://start.spring.io/starter.zip"
)

type Inputs struct {
	Type         string
	Language     string
	BootVersion  string
	BaseDir      string
	GroupId      string
	ArtifactId   string
	Name         string
	Description  string
	PackageName  string
	Packaging    string
	JavaVersion  string
	Dependencies string
}

func (in Inputs) Run() {
	log.Println("Starting scaffold generation...")
	fmt.Printf("Name: %v\n", in.Name)
	fmt.Printf("Description: %v\n", in.Description)

	zipFile, err := downloadFile(in)

	if err != nil {
		log.Fatal("Failed to download starter project\n", err)
	}

	if err := unzipFile(zipFile); err != nil {
		log.Fatal("Failed to Unzip file", err)
	}

	log.Println("Finished scaffold generation")
}

func downloadFile(inputs Inputs) (string, error) {
	log.Println("Starting download project.")

	req, err := http.NewRequest("GET", starterURL, nil)
	if err != nil {
		return "", err
	}

	q := req.URL.Query()
	q.Add("type", inputs.Type)
	q.Add("language", inputs.Language)
	q.Add("bootVersion", inputs.BootVersion)
	q.Add("baseDir", inputs.ArtifactId)
	q.Add("artifactId", inputs.ArtifactId)
	q.Add("groupId", inputs.GroupId)
	q.Add("name", inputs.ArtifactId)
	q.Add("description", inputs.Description)
	q.Add("packageName", inputs.GroupId+"."+inputs.ArtifactId)
	q.Add("packaging", inputs.Packaging)
	q.Add("javaVersion", inputs.JavaVersion)
	q.Add("dependencies", inputs.Dependencies)
	req.URL.RawQuery = q.Encode()
	log.Println(req.URL)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return "", err
	}

	//defer resp.Body.Close()
	defer func() {
		if err := resp.Body.Close(); err != nil {
			fmt.Sprintf("error closing http request to: %v, detail: %s", starterURL, err)
		}
	}()

	if resp.StatusCode != 200 {
		err := fmt.Errorf("Invalid parameters ou dependencies! Response Status Code: %s", resp.Status)
		return "", err
	}

	prjfile := fmt.Sprintf("%s.zip", inputs.Name)
	out, err := os.Create(prjfile)

	if err != nil {
		return "", err
	}
	//defer out.Close()
	defer func() {
		if err := out.Close(); err != nil {
			fmt.Sprintf("error closing %q, detail: %s", prjfile, err)
		}
	}()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}
	log.Println("Download done.")
	return prjfile, nil
}

func unzipFile(filename string) error {
	log.Println("Unzip files...")
	destFolder := strings.Replace(filename, ".zip", "", 1)
	fileutil.CreateIfNotExists(destFolder, 0755)
	err := fileutil.Unzip(filename, destFolder)
	if err != nil {
		return err
	}
	err = fileutil.RemoveFile(filename)
	if err != nil {
		return err
	}
	log.Println("Unzip done.")
	return nil
}
