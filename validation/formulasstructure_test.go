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

func TestFormulasContent(t *testing.T) {

	root := "../aws"
	rules := createRules()
	requiredFromRoot := getFullContentFromFormula(root)

	for key, value := range requiredFromRoot {

		fmt.Println(strings.ReplaceAll(key, "/", " "))

		files, _ := returnFilesFromFolder(key)

		fmt.Println(rules[value].Language)

		errFiles := diff(rules[value].Files, files)
		if len(errFiles) >= 1 {
			t.Errorf("Missing an files = %v", errFiles)
			return
		}
		errFolder := diff(rules[value].Folders, files)
		if len(errFolder) >= 1 {
			t.Errorf("Missing an folders = %v", errFolder)
			return
		}

	}

}

func getFullContentFromFormula(root string) map[string]string{
	required := map[string]string{}

	files, err := returnFilesFromFolder(root)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.Contains(file, "src/main.") {
			fileX :=  strings.Split(file, "/src/main.")
			required[fileX[0]] = fileX[1]
		}
	}
	return required
}

func returnFilesFromFolder(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	return files, err
}

func createRules() map[string]Rule {
	return map[string]Rule{
		"go" : {
			Language:   "Golang",
			Extensions: "go",
			Files:      []string{"go.mod", "main.go", "README.md", "config.json"},
			Folders:    []string{"src", "pkg"},
		},
		"java" : {
			Language: "Java",
			Extensions: "java",
			Files:      []string{"main.java", "README.md", "config.json"},
			Folders:    []string{"src", "pkg"},
		},
		"sh" : {
			Language: "shell script",
			Extensions: "sh",
			Files:      []string{"main.sh", "README.md", "config.json", "Makefile"},
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

