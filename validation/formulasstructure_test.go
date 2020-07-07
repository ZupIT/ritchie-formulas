package validation

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)
type Rule struct {
	Language string
	Extensions string
	Files    []string
	Folders  []string
}

func TestHandler(t *testing.T) {

	root := "../circleci"

	rules := createRules()

	requiredFromRoot := getFullContentFromFormula(root)

	for _, rule := range rules {
		if strings.HasSuffix(requiredFromRoot["main"], rule.Extensions){

			fmt.Println(strings.ReplaceAll(requiredFromRoot["root"], "/", " "))
			fmt.Println(rule.Language)

			_, files := returnFilesFromFolder(requiredFromRoot["root"])

			errFiles := diff(rule.Files, files)
			if len(errFiles) >= 1 {
				t.Errorf("Missing an files = %v", errFiles)
				return
			}
			errFolder := diff(rule.Folders, files)
			if len(errFolder) >= 1 {
				t.Errorf("Missing an folders = %v", errFolder )
				return
			}

		}
	}

}

func getFullContentFromFormula(root string) map[string]string{
	required := map[string]string{}

	err, files := returnFilesFromFolder(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.HasSuffix(file, "src") {
			required["root"] = strings.ReplaceAll(file, "/src", "")
		}
		if strings.Contains(file, "main") {
			required["main"] = file
		}
	}
	return required
}

func returnFilesFromFolder(root string) (error, []string) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	return err, files
}

func createRules() []Rule {
	return []Rule {
		{
			Language:   "go",
			Extensions: "go",
			Files:      []string{"go.mod", "main.go", "README.md", "config.json"},
			Folders:    []string{"src", "pkg"},
		},
		{
			Language: "java",
			Extensions: "java",
			Files:      []string{"main.java", "README.md", "config.json"},
			Folders:    []string{"src", "pkg"},
		}}
}

func diff(a, b []string) []string {
	if len(a) < 1  || len(b) < 1{
		return []string{"Empty Array"}
	}
	iFound := map[string]bool{}
	for _, itemA := range a {
		for _, itemB := range b {
			if strings.HasSuffix(itemB, itemA){
				iFound[itemA] = true
				break
			} else {
				iFound[itemA] = false
			}
		}

	}

	var failed []string
	for key, value := range iFound {
		if value {
			delete(iFound, key)
		} else {
			failed = append(failed, key)
		}
	}

	return failed
}

