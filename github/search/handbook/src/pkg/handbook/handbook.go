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

	"github.com/gookit/color"
)

const (
	zupGitSearchUrl = "https://api.github.com/search/code?q={{WORD}}+in:file+repo:zupit/{{REPOSITORY}}"
)

type Archive struct {
	Name    string `json:"name"`
	Url     string `json:"url"`
	Path    string `json:"path"`
	Content string `json:"content"`
}

type ResultSearch struct {
	TotalCount int       `json:"total_count"`
	Items      []Archive `json:"items"`
}

type Inputs struct {
	RepositoryName string
	GitUser        string
	GitToken       string
}

func (in Inputs) Run() {
	log.Println("Handbook Search Code Starter ...")
	repository := readRepository()
	word := readWord()

	url := strings.ReplaceAll(zupGitSearchUrl, "{{REPOSITORY}}", repository)
	url = strings.ReplaceAll(url, "{{WORD}}", word)

	resultSearch := in.searchRepository(url)
	if resultSearch.TotalCount <= 0 {
		log.Fatal("Not Found!")
		return
	}
	str := fmt.Sprintf("Found %d in file(s):", resultSearch.TotalCount)

	strSelect, _ := prompt.List(str, resultSearchToString(resultSearch))
	in.showContent(strSelect, word, resultSearch)

}

func (in Inputs) showContent(strSelect, word string, resultSearch ResultSearch) {
	a := stringToArchive(strSelect, resultSearch)

	req, err := http.NewRequest("GET", a.Url, nil)
	if err != nil {
		log.Fatal("Error to scan Repository Request: ", err)
	}
	req.SetBasicAuth(in.GitUser, in.GitToken)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("Error process scan Repository: ", err)
	}
	defer resp.Body.Close()

	var ar Archive
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &ar)
	if err != nil {
		log.Fatal("Error proccess convert json to struct:", err)
	}
	c := decodeContent(ar.Content)
	color.Success.Println(c)
}

func stringToArchive(strSelect string, resultSearch ResultSearch) Archive {
	for _, r := range resultSearch.Items {
		if r.Path == strSelect {
			return r
		}
	}
	log.Fatal("Error, str not found in resultSearch")
	return Archive{}
}

func decodeContent(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		log.Println("error:", err)
		return ""
	}
	return string(data)
}

func resultSearchToString(rs ResultSearch) []string {
	var str []string

	for _, x := range rs.Items {
		str = append(str, x.Path)
	}

	return str
}

func (in Inputs) searchRepository(url string) ResultSearch {

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

	var rs ResultSearch
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(bodyBytes, &rs)
	if err != nil {
		log.Fatal("Error proccess convert json to struct:", err)
	}
	return rs
}

func readRepository() string {
	repository, err := prompt.String("Type name of application repository: ", false)
	if err != nil {
		log.Fatal(err)
	}
	return repository
}

func readWord() string {
	repository, err := prompt.String("Type word of search: ", false)
	if err != nil {
		log.Fatal(err)
	}
	return repository
}
