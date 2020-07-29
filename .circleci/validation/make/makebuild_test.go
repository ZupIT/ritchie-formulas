package make

import (
	"fmt"
	"strings"
	"testing"
)

const root = "../../../"

func TestBuild(t *testing.T) {
	dirs := getFullContentFromFormula(root)

	for dir, language := range dirs {
		if language == "go" {
			err := validateMakefile(dir)

			if err != nil {
				t.Errorf("%s", err)
			}

			path := strings.ReplaceAll(dir, "/", " ")
			fmt.Printf("Build success: %s\n", path)
		}
	}
}
