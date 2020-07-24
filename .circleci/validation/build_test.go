package validation

import (
	"fmt"
	"testing"
)

var root = "../../"

func TestBuild(t *testing.T) {
	dirs := getFullContentFromFormula(root)

	for dir, language := range dirs {
		if language == "go" {
			message, err := validateMakefile(dir)

			if err {
				t.Errorf(message)
			}

			fmt.Println(message)
		}
	}
}
