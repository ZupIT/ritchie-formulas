package structure

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func check(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}

func getFullContentFromFormula(root string) map[string]string {
	required := map[string]string{}

	err := os.Chdir(root)
	check(err)

	pwd, err := os.Getwd()
	check(err)

	files, err := returnFilesFromFolder(pwd)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if strings.Contains(file, "src/main.") && !strings.Contains(file, "templates") {
			fileX := strings.Split(file, "/src/main.")
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

func diff(a, b []string) []string {
	if len(a) < 1 || len(b) < 1 {
		return []string{"Empty Array"}
	}
	iFound := map[string]bool{}

	for _, itemA := range a {
		for _, itemB := range b {
			if strings.HasSuffix(itemB, itemA) {
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
