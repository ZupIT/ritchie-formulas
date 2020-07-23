/* THIS IS A DRAFT
 *
 * Steps:
 *  1. [X] Open formula folder
 *  2. [X] Check makefile exist
 *  3. [ ] Run make build
 *  4. [ ] Capture command return
 */
package validation

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

var root = "../"

func TestBuild(t *testing.T) {
	dirs := getFullContentFromFormula(root)

	for dir, language := range dirs {
		if language == "go" {
			fmt.Printf("Diretório: %s\nLinguagem: %s\nTem Makefile? ", dir, language)
			validateMakefile(dir)
			fmt.Print("\n")
		}
	}
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
			fmt.Println("Sim")
			runBuild(path)
		}
	}
}

func runBuild(path string) {
	var cmd *exec.Cmd
	var stderr bytes.Buffer

	cmd = exec.Command("bash", "-c", "cd", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if stderr.Bytes() != nil {
			errMsg := fmt.Sprintf("Build error: \n%s \n%s", stderr.String(), err)

			fmt.Println(errMsg)
		}
	}

	fmt.Println(cmd)
}
