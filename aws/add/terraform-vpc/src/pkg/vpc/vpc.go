package vpc

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"text/template"
	"vpc/pkg/tpl"

	"github.com/ZupIT/ritchie-cli/pkg/file/fileutil"
	"github.com/fatih/color"
	"github.com/hashicorp/terraform/configs"
)

const (
	projectFile = ".scaffold"
	maintfFile  = "src/main.tf"
	variableQA  = "src/variables/qa.tfvars"
)

type Inputs struct {
	Region       string
	VPCName      string
	VPCCIDR      string
	VPCAZS       string
	CustomerName string
	PWD          string
}

func Run(in Inputs) {
	cdir := in.PWD

	in.checkIfProjectExist()

	if !in.moduleExist() {
		//main.tf
		mainf, err := os.OpenFile(path.Join(cdir, maintfFile), os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			color.Yellow(fmt.Sprintf("error openning main.tf, detail: %q", err))
			os.Exit(1)
		}
		defer mainf.Close()
		if _, err := mainf.WriteString(tpl.Maintf); err != nil {
			color.Red(fmt.Sprintf("error writing main.tf, detail: %q", err))
			os.Exit(1)
		}

		// variables
		in.parseAZS()
		t := template.Must(template.New("Var").Parse(tpl.Variable))
		varf := path.Join(cdir, variableQA)
		vfile, err := os.OpenFile(varf, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			color.Red(fmt.Sprintf("error openning %q, detail: %q", varf, err))
			os.Exit(1)
		}
		defer vfile.Close()
		err = t.Execute(vfile, in)
		if err != nil {
			color.Red(fmt.Sprintf("error writing %q, detail: %q", varf, err))
			os.Exit(1)
		}
	}

	fmt.Println()
	color.Green(fmt.Sprintln("vpc module configured successfully"))
	color.Green(fmt.Sprintln("go to the src dir and run [ENVIRONMENT=qa make plan] to check the terraform plan"))
}

func (in Inputs) moduleExist() bool {
	parser := configs.NewParser(nil)
	cfg, diags := parser.LoadConfigFile(path.Join(in.PWD, maintfFile))
	if len(diags) != 0 {
		color.Red("unexpected diagnostics")
		for _, diag := range diags {
			color.Red(fmt.Sprintf("- %s", diag))
		}
		os.Exit(1)
	}

	for _, m := range cfg.ModuleCalls {
		if m.Name == "vpc" {
			return true
		}
	}
	return false
}

func (in *Inputs) parseAZS() {
	strings.Replace(in.VPCAZS, " ", "", -1)
	ss := strings.Split(in.VPCAZS, ",")
	siz := len(ss) - 1
	var azs string
	for i, s := range ss {
		sq := strconv.Quote(s)
		if i < siz {
			azs += sq + ","
		} else {
			azs += sq
		}
	}
	in.VPCAZS = azs
}

func (in Inputs) checkIfProjectExist() {
	if !fileutil.Exists(path.Join(in.PWD, projectFile)) {
		color.Red("seems that your current dir isn't a terraform project.")
		color.Red("you can create one running [rit aws generate terraform-project]")
		os.Exit(1)
	}
}
