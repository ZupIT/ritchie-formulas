/* THIS IS A DRAFT
 *
 * Steps:
 *  1. [X] Open formula folder
 *  2. [X] Check makefile exist
 *  3. [X] Run make build
 *  4. [X] Capture command return
 *  5. [ ] Run build.bat (for windows)
 *  6. [ ] Run make test
 *  7. [ ] Parallel Tests
 */
package validation

import (
	"testing"
)

var root = "../../"

func TestBuild(t *testing.T) {
	dirs := getFullContentFromFormula(root)

	for dir, language := range dirs {
		if language == "go" {
			validateMakefile(dir)
		}
	}
}
