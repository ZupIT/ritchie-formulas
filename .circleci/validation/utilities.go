package validation

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

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
		if strings.Contains(file, "src/main.") && !strings.HasPrefix(file, "../../templates") {
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

func check(err error) {
	if err != nil {
		fmt.Printf("Erro: %v\n", err)
	}
}

func validateMakefile(path string) {
	files, err := ioutil.ReadDir(path)
	check(err)

	for _, file := range files {
		if file.Name() == "Makefile" {
			runBuild(path)
		}
	}
}

func runBuild(path string) {
	var cmd *exec.Cmd
	var stderr, stdout bytes.Buffer

	err := os.Chdir(path)
	check(err)

	cmd = exec.Command("make", "build")
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout

	if err := cmd.Run(); err != nil {
		if stderr.Bytes() != nil {
			fmt.Sprintf("Build failed: \n%s \n%s", stderr.String(), err)
		}
	}

	if stderr.String() != "" {
		fmt.Printf("Build failed: %s\n", stderr.String())
		os.Exit(1)
	}

	fmt.Printf("Build success: %s\n", path)
}
