package eks

import (
	"fmt"
	"html/template"
	"os"
	"path"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"

	"eks/pkg/tpl"

	"github.com/ZupIT/ritchie-cli/pkg/file/fileutil"
	"github.com/fatih/color"
)

const (
	projectFile   = ".scaffold"
	maintfFile    = "src/main.tf"
	mainEKStfFile = "pkg/tpl/main.tf"
	dnsZoneModule = "src/modules/dns_zone"
	hemlDeps      = "src/modules/helm_deps"
	iamK8SModule  = "src/modules/iam_k8s"
	variableQA    = "src/variables/qa.tfvars"
)

const (
	terraform = "terraform"
)

type Inputs struct {
	ClusterName string
	DomainName  string
	PWD         string
}

func Run(in Inputs) {
	in.checkIfProjectExist()

	in.mergeMain()

	in.addVariables()

	in.addDNSZone()

	in.addHelmDeps()

	in.addIAMK8S()

	fmt.Println()
	color.Green(fmt.Sprintln("eks module configured successfully"))
	color.Green(fmt.Sprintln("go to the src dir and run [ENVIRONMENT=qa make plan] to check the terraform plan"))
}

func (in Inputs) addIAMK8S() {
	// iam_k8s
	iamdir := path.Join(in.PWD, iamK8SModule)
	if err := fileutil.CreateDirIfNotExists(iamdir, 0755); err != nil {
		color.Red(fmt.Sprintf("error creating dir %q, detail: %q", iamdir, err))
		os.Exit(1)
	}

	iammain := path.Join(iamdir, "main.tf")
	if err := fileutil.CreateFileIfNotExist(iammain, []byte(tpl.IAMK8SMaintf)); err != nil {
		color.Red(fmt.Sprintf("error creating file %q, detail: %q", iammain, err))
		os.Exit(1)
	}

	iamvar := path.Join(iamdir, "variables.tf")
	if err := fileutil.CreateFileIfNotExist(iamvar, []byte(tpl.IAMK8SVariablestf)); err != nil {
		color.Red(fmt.Sprintf("error creating file %q, detail: %q", iamvar, err))
		os.Exit(1)
	}
}

func (in Inputs) addHelmDeps() {
	// helm_deps
	helmdir := path.Join(in.PWD, hemlDeps)
	if err := fileutil.CreateDirIfNotExists(helmdir, 0755); err != nil {
		color.Red(fmt.Sprintf("error creating dir %q, detail: %q", helmdir, err))
		os.Exit(1)
	}

	helmmain := path.Join(helmdir, "main.tf")
	if err := fileutil.CreateFileIfNotExist(helmmain, []byte(tpl.HelmMaintf)); err != nil {
		color.Red(fmt.Sprintf("error creating file %q, detail: %q", helmmain, err))
		os.Exit(1)
	}

	helmvar := path.Join(helmdir, "variables.tf")
	if err := fileutil.CreateFileIfNotExist(helmvar, []byte(tpl.HelmVariablestf)); err != nil {
		color.Red(fmt.Sprintf("error creating file %q, detail: %q", helmvar, err))
		os.Exit(1)
	}
}

func (in Inputs) addDNSZone() {
	// dns_zone
	dzdir := path.Join(in.PWD, dnsZoneModule)
	if err := fileutil.CreateDirIfNotExists(dzdir, 0755); err != nil {
		color.Red(fmt.Sprintf("error creating dir %q, detail: %q", dzdir, err))
		os.Exit(1)
	}

	dzmain := path.Join(dzdir, "main.tf")
	if err := fileutil.CreateFileIfNotExist(dzmain, []byte(tpl.DnsZoneMaintf)); err != nil {
		color.Red(fmt.Sprintf("error creating file %q, detail: %q", dzmain, err))
		os.Exit(1)
	}

	dzout := path.Join(dzdir, "outputs.tf")
	if err := fileutil.CreateFileIfNotExist(dzout, []byte(tpl.DnsZoneOutputstf)); err != nil {
		color.Red(fmt.Sprintf("error creating file %q, detail: %q", dzout, err))
		os.Exit(1)
	}

	dzvar := path.Join(dzdir, "variables.tf")
	if err := fileutil.CreateFileIfNotExist(dzvar, []byte(tpl.DnsZoneVariablestf)); err != nil {
		color.Red(fmt.Sprintf("error creating file %q, detail: %q", dzvar, err))
		os.Exit(1)
	}
}

func (in Inputs) addVariables() {
	// variables
	t := template.Must(template.New("Var").Parse(tpl.Variable))
	varf := path.Join(in.PWD, variableQA)
	vf, err := os.OpenFile(varf, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		color.Red(fmt.Sprintf("error openning %q, detail: %q", varf, err))
		os.Exit(1)
	}
	defer func() {
		if err := vf.Close(); err != nil {
			color.Yellow(fmt.Sprintf("error closing %q, detail: %s", varf, err))
		}
	}()

	err = t.Execute(vf, in)
	if err != nil {
		color.Red(fmt.Sprintf("error writing %q, detail: %q", varf, err))
		os.Exit(1)
	}
}

func (in Inputs) mergeMain() {
	// main.tf current
	mfile := path.Join(in.PWD, maintfFile)
	mb, _ := fileutil.ReadFile(mfile)
	mcfg, diags := hclwrite.ParseConfig(mb, mfile, hcl.InitialPos)
	checkDiagnostics(diags)

	// main.tk eks
	dir, _ := os.Getwd()
	efile := path.Join(dir, mainEKStfFile)
	eb, _ := fileutil.ReadFile(efile)
	ecfg, diags := hclwrite.ParseConfig(eb, efile, hcl.InitialPos)
	checkDiagnostics(diags)

	var reqblk *hclwrite.Block
	mbody := mcfg.Body()
	for _, block := range mbody.Blocks() {
		if block.Type() == terraform {
			for _, tf := range block.Body().Blocks() {
				if tf.Type() == "required_providers" {
					reqblk = tf
					break
				}
			}
		}
	}

	ebody := ecfg.Body()
	for _, block := range ebody.Blocks() {
		if block.Type() == terraform {
			for _, tf := range block.Body().Blocks() {
				if tf.Type() == "required_providers" {
					for n, a := range tf.Body().Attributes() {
						var tokens hclwrite.Tokens
						tokens = a.BuildTokens(tokens)
						reqblk.Body().SetAttributeRaw(n, tokens[2:len(tokens)-1])
					}
				}
			}
		}
	}

	fileutil.WriteFile(mfile, mcfg.Bytes())

	// main.tf others
	mf, err := os.OpenFile(mfile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		color.Red(fmt.Sprintf("error openning %q, detail: %q", mfile, err))
		os.Exit(1)
	}
	defer func() {
		if err := mf.Close(); err != nil {
			color.Yellow(fmt.Sprintf("error closing %q, detail: %s", mfile, err))
		}
	}()
	_, err = mf.Write([]byte(tpl.Maintf))
	if err != nil {
		color.Red(fmt.Sprintf("error appending file %q, detail: %q", mfile, err))
		os.Exit(1)
	}
}

func (in Inputs) checkIfProjectExist() {
	if !fileutil.Exists(path.Join(in.PWD, projectFile)) {
		color.Red("seems that your current dir isn't a terraform project.")
		color.Red("you can create one running [rit aws generate terraform-project]")
		os.Exit(1)
	}
}

func checkDiagnostics(diags hcl.Diagnostics) {
	if len(diags) != 0 {
		color.Red("unexpected diagnostics")
		for _, diag := range diags {
			color.Red(fmt.Sprintf("- %s", diag))
		}
		os.Exit(1)
	}
}
