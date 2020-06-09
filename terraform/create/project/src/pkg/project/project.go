package project

import (
	"fmt"
	"io/ioutil"
	"os"

	"project/pkg/tpl"

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

func (in Input) Run() {
	//log the action
	if err := os.MkdirAll(in.RepoPath(), 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create dir: '%s', error: '%s'", in.RepoPath(), err.Error()))
		os.Exit(1)
	}

	readme := fmt.Sprintf(readmeFormat, in.RepoPath())
	if err := ioutil.WriteFile(readme, []byte(tpl.Readme), 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create readme: '%s', error: '%s'", readme, err.Error()))
		os.Exit(1)
	}

	gitignore := fmt.Sprintf(gitignoreFormat, in.RepoPath())
	if err := ioutil.WriteFile(gitignore, []byte(tpl.GitIgnore), 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create gitignore: '%s', error: '%s'", readme, err.Error()))
		os.Exit(1)
	}

	jenkinsfile := fmt.Sprintf(jenkinsfileFormat, in.RepoPath())
	if err := ioutil.WriteFile(jenkinsfile, []byte(tpl.Jenkinsfile), 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create jenkinsfile: '%s', error: '%s'", readme, err.Error()))
		os.Exit(1)
	}

	src := fmt.Sprintf(srcDir, in.RepoPath())
	if err := os.MkdirAll(src, 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create dir: '%s', error: '%s'", src, err.Error()))
		os.Exit(1)
	}

	maintf := fmt.Sprintf(mainFormat, in.RepoPath())
	if err := ioutil.WriteFile(maintf, []byte(tpl.Maintf), 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create main.tf: '%s', error: '%s'", maintf, err.Error()))
		os.Exit(1)
	}

	makefile := fmt.Sprintf(makefileFormat, in.RepoPath())
	if err := ioutil.WriteFile(makefile, []byte(tpl.Makefile), 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create Makefile: '%s', error: '%s'", maintf, err.Error()))
		os.Exit(1)
	}

	modules := fmt.Sprintf(modulesDir, in.RepoPath())
	if err := os.MkdirAll(modules, 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create dir: '%s', error: '%s'", modules, err.Error()))
		os.Exit(1)
	}

	variables := fmt.Sprintf(variablesDir, in.RepoPath())
	if err := os.MkdirAll(variables, 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create dir: '%s', error: '%s'", variables, err.Error()))
		os.Exit(1)
	}

	commonvar := fmt.Sprintf(varFilesFormat, in.RepoPath(), "common")
	if err := ioutil.WriteFile(commonvar, []byte(tpl.Variable), 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create variable file: '%s', error: '%s'", commonvar, err.Error()))
		os.Exit(1)
	}

	prodvar := fmt.Sprintf(varFilesFormat, in.RepoPath(), "prod")
	if err := ioutil.WriteFile(prodvar, []byte(tpl.Variable), 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create variable file: '%s', error: '%s'", prodvar, err.Error()))
		os.Exit(1)
	}

	qavar := fmt.Sprintf(varFilesFormat, in.RepoPath(), "qa")
	if err := ioutil.WriteFile(qavar, []byte(tpl.Variable), 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create variable file: '%s', error: '%s'", qavar, err.Error()))
		os.Exit(1)
	}

	stgvar := fmt.Sprintf(varFilesFormat, in.RepoPath(), "stg")
	if err := ioutil.WriteFile(stgvar, []byte(tpl.Variable), 0755); err != nil {
		color.Red(fmt.Sprintf("failed to create variable file: '%s', error: '%s'", stgvar, err.Error()))
		os.Exit(1)
	}

	//log verdinho sucesso com info

}
