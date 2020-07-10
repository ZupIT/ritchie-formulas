package validation

import (
	"fmt"
	"strings"
	"testing"
)

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
