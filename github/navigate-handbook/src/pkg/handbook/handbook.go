package handbook

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"handbook/pkg/prompt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	zupGitListUrl = "https://api.github.com/repos/ZupIT/{{REPOSITORY}}/contents/"
)

type Archive struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Content string `json:"content"`
}

type Inputs struct {
	GitUser        string
	GitToken       string
}

func (in Inputs) Run() {
	log.Println("Handbook Navigate Starter ...")

	repository := readRepository()
	url := strings.ReplaceAll(zupGitListUrl, "{{REPOSITORY}}", repository)
	archives := in.scanRepository(url)

	strSelect, _ := prompt.List("Escolha", archivesToString(archives))
	in.navigateGit(archives,strSelect,url)
}

func (in Inputs)navigateGit(archives []Archive,strSelect ,url  string ){
	x := verifyTypeFile(archives,strSelect)

	if x {
		url = fmt.Sprint(url,"/"+strSelect)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal("Error to git Repository Request: ", err)
		}
		req.SetBasicAuth(in.GitUser, in.GitToken)
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Fatal("Error process git Repository: ", err)
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		var a Archive
		err = json.Unmarshal(bodyBytes, &a)
		if err != nil {
		log.Fatal("Error proccess convert json to struct:", err)
		}

		log.Println(decodeContent(a.Content))
		return
	}
	url = fmt.Sprint(url,"/"+strSelect)

	archives = in.scanRepository(url)
	strSelect, _ = prompt.List("Escolha", archivesToString(archives))
	in.navigateGit(archives, strSelect, url)
}

func decodeContent(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Println("error:", err)
		return ""
	}
	return string(data)
}

func archivesToString(archives []Archive) []string {
	var str []string

	for _, a := range archives {
		str = append(str, a.Name)
	}
	return str
}

func readRepository() string {
	repository, err := prompt.String("Type name of application repository: ", false)
	if err != nil {
		log.Fatal(err)
	}
	return repository
}

func (in Inputs) scanRepository(url string) []Archive {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Error to scan Repository Request: ", err)
	}
	req.SetBasicAuth(in.GitUser, in.GitToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error process scan Repository: ", err)
	}
	defer resp.Body.Close()

	var archives []Archive
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(bodyBytes, &archives)
	if err != nil {
		log.Fatal("Error proccess convert json to struct:", err)
	}
	return archives
}

func verifyTypeFile(archives []Archive, str string) bool{

	for _,a := range archives  {
		if a.Name == str {
			switch a.Type {
			case "file":
				return true
			case "dir":
				return false
			default:
				log.Fatal("Type GitHub is not valid.")
			}
		}
	}
	return false
}