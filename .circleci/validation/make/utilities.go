package make

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
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
		if strings.Contains(file, "src/main.") && !strings.HasPrefix(file, "templates") {
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

func validateMakefile(path string) error {
	files, err := ioutil.ReadDir(path)
	check(err)

	for _, file := range files {
		if file.Name() == "Makefile" {
			err := runBuild(path)
			if err != nil {
				path = strings.ReplaceAll(path, "/", " ")
				errMsg := fmt.Sprintf("Build failed: %s\n%s", path, err)

				return errors.New(errMsg)
			}
		}
	}

	return nil
}

func runBuild(path string) error {
	var cmd *exec.Cmd
	var stderr bytes.Buffer

	if err := os.Chdir(path); err != nil {
		return err
	}

	cmd = exec.Command("make", "build")
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		if stderr.Bytes() != nil {
			errMsg := fmt.Sprintf("Build failed: \n%s \n%s", stderr.String(), err)

			return errors.New(errMsg)
		}

		return err
	}

	return nil
}
