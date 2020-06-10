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

	if err := CreateDirIfNotExists(in.RepoPath(), 0755); err != nil {
		in.rollback(err)
	}

	readme := fmt.Sprintf(readmeFormat, in.RepoPath())
	if err := CreateFileIfNotExist(readme, []byte(tpl.Readme)); err != nil {
		in.rollback(err)
	}

	gitignore := fmt.Sprintf(gitignoreFormat, in.RepoPath())
	if err := CreateFileIfNotExist(gitignore, []byte(tpl.GitIgnore)); err != nil {
		in.rollback(err)
	}

	jenkinsfile := fmt.Sprintf(jenkinsfileFormat, in.RepoPath())
	if err := CreateFileIfNotExist(jenkinsfile, []byte(tpl.Jenkinsfile)); err != nil {
		in.rollback(err)
	}

	src := fmt.Sprintf(srcDir, in.RepoPath())
	if err := CreateDirIfNotExists(src, 0755); err != nil {
		in.rollback(err)
	}

	maintf := fmt.Sprintf(mainFormat, in.RepoPath())
	if err := CreateFileIfNotExist(maintf, []byte(tpl.Maintf)); err != nil {
		in.rollback(err)
	}

	makefile := fmt.Sprintf(makefileFormat, in.RepoPath())
	if err := CreateFileIfNotExist(makefile, []byte(tpl.Makefile)); err != nil {
		in.rollback(err)
	}

	modules := fmt.Sprintf(modulesDir, in.RepoPath())
	if err := CreateDirIfNotExists(modules, 0755); err != nil {
		in.rollback(err)
	}

	variables := fmt.Sprintf(variablesDir, in.RepoPath())
	if err := CreateDirIfNotExists(variables, 0755); err != nil {
		in.rollback(err)
	}

	commonvar := fmt.Sprintf(varFilesFormat, in.RepoPath(), "common")
	if err := CreateFileIfNotExist(commonvar, []byte(tpl.Variable)); err != nil {
		in.rollback(err)
	}

	prodvar := fmt.Sprintf(varFilesFormat, in.RepoPath(), "prod")
	if err := CreateFileIfNotExist(prodvar, []byte(tpl.Variable)); err != nil {
		in.rollback(err)
	}

	qavar := fmt.Sprintf(varFilesFormat, in.RepoPath(), "qa")
	if err := CreateFileIfNotExist(qavar, []byte(tpl.Variable)); err != nil {
		in.rollback(err)
	}

	stgvar := fmt.Sprintf(varFilesFormat, in.RepoPath(), "stg")
	if err := CreateFileIfNotExist(stgvar, []byte(tpl.Variable)); err != nil {
		in.rollback(err)
	}

	color.Green(fmt.Sprintln("project created successfully."))
	color.Green(fmt.Sprintln("Location:", in.RepoPath()))
	color.Green(fmt.Sprintln("Run [rit terraform aws] and check the formulas module"))

}

func IsNotExist(name string) bool {
	_, err := os.Stat(name)
	return err != nil && os.IsNotExist(err)
}

func CreateDirIfNotExists(dir string, perm os.FileMode) error {
	if IsNotExist(dir) {
		if err := os.MkdirAll(dir, perm); err != nil {
			return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
		}
	}
	return nil
}

func CreateFileIfNotExist(file string, content []byte) error {
	if IsNotExist(file) {
		if err := ioutil.WriteFile(file, content, 0644); err != nil {
			return err
		}
	}
	return nil
}
