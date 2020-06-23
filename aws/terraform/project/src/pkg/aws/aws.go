package aws

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"project/pkg/tpl"

	"github.com/fatih/color"

	"github.com/ZupIT/ritchie-cli/pkg/file/fileutil"
)

const (
	dirFormat          = "%s/%s"
	scaffoldFormat     = "%s/.scaffold"
	readmeFormat       = "%s/README.md"
	gitignoreFormat    = "%s/.gitignore"
	jenkinsfileFormat  = "%s/Jenkinsfile"
	srcDir             = "%s/src"
	mainFormat         = "%s/src/main.tf"
	makefileFormat     = "%s/src/Makefile"
	modulesDir         = "%s/src/modules"
	templatesDir       = "%s/src/templates"
	variablesDir       = "%s/src/variables"
	varFilesFormat     = "%s/src/variables/%s.tfvars"
	QABackendtfFormat  = "%s/src/qa.tfbackend"
	CircleCIPath       = "files/circleci-pipeline"
	CircleCIConfigPath = "%s/.circleci/config.yml"
)

type Input struct {
	ProjectName     string
	ProjectLocation string
	BucketName      string
	BucketRegion    string
	PWD             string
}

func (in Input) Path() string {
	return fmt.Sprintf(dirFormat, in.ProjectLocation, in.ProjectName)
}

func (in Input) rollback(err error) {
	if err != nil {
		color.Red(fmt.Sprintf("failed to create project: %q, error: %q", in.Path(), err.Error()))
		if err := os.RemoveAll(in.Path()); err != nil {
			color.Red(fmt.Sprintf("failed to rollback: %q, error: %q", in.Path(), err.Error()))
		}
		os.Exit(1)
	}
}

func (in Input) Run() {

	if err := CreateDirIfNotExists(in.Path(), 0755); err != nil {
		in.rollback(err)
	}

	scaffold := fmt.Sprintf(scaffoldFormat, in.Path())
	if err := CreateFileIfNotExist(scaffold, []byte(tpl.Scaffold)); err != nil {
		in.rollback(err)
	}

	readme := fmt.Sprintf(readmeFormat, in.Path())
	if err := CreateFileIfNotExist(readme, []byte(tpl.Readme)); err != nil {
		in.rollback(err)
	}

	gitignore := fmt.Sprintf(gitignoreFormat, in.Path())
	if err := CreateFileIfNotExist(gitignore, []byte(tpl.GitIgnore)); err != nil {
		in.rollback(err)
	}

	jenkinsfile := fmt.Sprintf(jenkinsfileFormat, in.Path())
	if err := CreateFileIfNotExist(jenkinsfile, []byte(tpl.Jenkinsfile)); err != nil {
		in.rollback(err)
	}

	src := fmt.Sprintf(srcDir, in.Path())
	if err := CreateDirIfNotExists(src, 0755); err != nil {
		in.rollback(err)
	}

	maintf := fmt.Sprintf(mainFormat, in.Path())
	if err := CreateFileIfNotExist(maintf, []byte(tpl.Maintf)); err != nil {
		in.rollback(err)
	}

	makefile := fmt.Sprintf(makefileFormat, in.Path())
	if err := CreateFileIfNotExist(makefile, []byte(tpl.Makefile)); err != nil {
		in.rollback(err)
	}

	modules := fmt.Sprintf(modulesDir, in.Path())
	if err := CreateDirIfNotExists(modules, 0755); err != nil {
		in.rollback(err)
	}

	templates := fmt.Sprintf(templatesDir, in.Path())
	if err := CreateDirIfNotExists(templates, 0755); err != nil {
		in.rollback(err)
	}

	variables := fmt.Sprintf(variablesDir, in.Path())
	if err := CreateDirIfNotExists(variables, 0755); err != nil {
		in.rollback(err)
	}

	commonvar := fmt.Sprintf(varFilesFormat, in.Path(), "common")
	if err := CreateFileIfNotExist(commonvar, []byte(tpl.Variable)); err != nil {
		in.rollback(err)
	}

	prodvar := fmt.Sprintf(varFilesFormat, in.Path(), "prod")
	if err := CreateFileIfNotExist(prodvar, []byte(tpl.Variable)); err != nil {
		in.rollback(err)
	}

	qavar := fmt.Sprintf(varFilesFormat, in.Path(), "qa")
	if err := CreateFileIfNotExist(qavar, []byte(tpl.Variable)); err != nil {
		in.rollback(err)
	}

	stgvar := fmt.Sprintf(varFilesFormat, in.Path(), "stg")
	if err := CreateFileIfNotExist(stgvar, []byte(tpl.Variable)); err != nil {
		in.rollback(err)
	}

	//qa.backendtf
	backendtf := fmt.Sprintf(QABackendtfFormat, in.Path())
	if err := CreateFileIfNotExist(backendtf, []byte("")); err != nil {
		in.rollback(err)
	}

	t := template.Must(template.New("QABackendtf").Parse(tpl.QABackendtf))
	bfile, err := os.OpenFile(backendtf, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		in.rollback(err)
	}
	defer bfile.Close()
	err = t.Execute(bfile, in)
	if err != nil {
		in.rollback(err)
	}

	color.Blue("Copying circleci pipeline files")

	circleDir := path.Join(in.PWD, in.ProjectName, ".circleci")

	if err := fileutil.CreateDirIfNotExists(circleDir, 0755); err != nil {
		in.rollback(err)
	}

	if err = fileutil.CopyDirectory(CircleCIPath, circleDir); err != nil {
		in.rollback(err)
	}

	circleciConfig := fmt.Sprintf(CircleCIConfigPath, in.Path())
	if err := CreateFileIfNotExist(circleciConfig, []byte("")); err != nil {
		in.rollback(err)
	}

	t = template.Must(template.New("circleciConfig").Parse(tpl.Circleciconfig))
	bfile, err = os.OpenFile(circleciConfig, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		in.rollback(err)
	}
	defer bfile.Close()
	err = t.Execute(bfile, in)
	if err != nil {
		in.rollback(err)
	}

	color.Green(fmt.Sprintln("project created successfully"))
	color.Green(fmt.Sprintln("location:", in.Path()))
	color.Green(fmt.Sprintln("go to the location and run [rit aws add] and check the options for your project"))

}

func IsNotExist(name string) bool {
	_, err := os.Stat(name)
	return err != nil && os.IsNotExist(err)
}

func CreateDirIfNotExists(dir string, perm os.FileMode) error {
	if IsNotExist(dir) {
		if err := os.MkdirAll(dir, perm); err != nil {
			return fmt.Errorf("failed to create directory: %q, error: %q", dir, err)
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
