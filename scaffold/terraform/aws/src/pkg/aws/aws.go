package aws

import (
	"fmt"
	"io/ioutil"
	"os"

	"aws/pkg/tpl"

	"github.com/fatih/color"
)

const (
	dirFormat         = "%s/%s"
	readmeFormat      = "%s/README.md"
	gitignoreFormat   = "%s/.gitignore"
	jenkinsfileFormat = "%s/Jenkinsfile"
	srcDir            = "%s/src"
	mainFormat        = "%s/src/main.tf"
	makefileFormat    = "%s/src/Makefile"
	modulesDir        = "%s/src/modules"
	variablesDir      = "%s/src/variables"
	varFilesFormat    = "%s/src/variables/%s.tfvars"
)

type Input struct {
	ProjectName string
	ProjectPath string
}

func (in Input) RepoPath() string {
	return fmt.Sprintf(dirFormat, in.ProjectPath, in.ProjectName)
}

func (in Input) rollback(err error) {
	if err != nil {
		color.Red(fmt.Sprintf("failed to create project: '%s', error: '%s'", in.RepoPath(), err.Error()))
		if err := os.RemoveAll(in.RepoPath()); err != nil {
			color.Red(fmt.Sprintf("failed to rollback: '%s', error: '%s'", in.RepoPath(), err.Error()))
		}
		os.Exit(1)
	}
}

func (in Input) Run() {
	if err := os.MkdirAll(in.RepoPath(), 0755); err != nil {
		in.rollback(err)
	}

	readme := fmt.Sprintf(readmeFormat, in.RepoPath())
	if err := ioutil.WriteFile(readme, []byte(tpl.Readme), 0755); err != nil {
		in.rollback(err)
	}

	gitignore := fmt.Sprintf(gitignoreFormat, in.RepoPath())
	if err := ioutil.WriteFile(gitignore, []byte(tpl.GitIgnore), 0755); err != nil {
		in.rollback(err)
	}

	jenkinsfile := fmt.Sprintf(jenkinsfileFormat, in.RepoPath())
	if err := ioutil.WriteFile(jenkinsfile, []byte(tpl.Jenkinsfile), 0755); err != nil {
		in.rollback(err)
	}

	src := fmt.Sprintf(srcDir, in.RepoPath())
	if err := os.MkdirAll(src, 0755); err != nil {
		in.rollback(err)
	}

	maintf := fmt.Sprintf(mainFormat, in.RepoPath())
	if err := ioutil.WriteFile(maintf, []byte(tpl.Maintf), 0755); err != nil {
		in.rollback(err)
	}

	makefile := fmt.Sprintf(makefileFormat, in.RepoPath())
	if err := ioutil.WriteFile(makefile, []byte(tpl.Makefile), 0755); err != nil {
		in.rollback(err)
	}

	modules := fmt.Sprintf(modulesDir, in.RepoPath())
	if err := os.MkdirAll(modules, 0755); err != nil {
		in.rollback(err)
	}

	variables := fmt.Sprintf(variablesDir, in.RepoPath())
	if err := os.MkdirAll(variables, 0755); err != nil {
		in.rollback(err)
	}

	commonvar := fmt.Sprintf(varFilesFormat, in.RepoPath(), "common")
	if err := ioutil.WriteFile(commonvar, []byte(tpl.Variable), 0755); err != nil {
		in.rollback(err)
	}

	prodvar := fmt.Sprintf(varFilesFormat, in.RepoPath(), "prod")
	if err := ioutil.WriteFile(prodvar, []byte(tpl.Variable), 0755); err != nil {
		in.rollback(err)
	}

	qavar := fmt.Sprintf(varFilesFormat, in.RepoPath(), "qa")
	if err := ioutil.WriteFile(qavar, []byte(tpl.Variable), 0755); err != nil {
		in.rollback(err)
	}

	stgvar := fmt.Sprintf(varFilesFormat, in.RepoPath(), "stg")
	if err := ioutil.WriteFile(stgvar, []byte(tpl.Variable), 0755); err != nil {
		in.rollback(err)
	}

	color.Green(fmt.Sprintln("project created successfully."))
	color.Green(fmt.Sprintln("Location:", in.RepoPath()))
	color.Green(fmt.Sprintln("Run [rit terraform aws] and check the formulas module"))

}
