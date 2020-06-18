package eks

import (
	"fmt"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"html/template"
	"os"
	"path"

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

type Inputs struct {
	ClusterName string
	DomainName  string
	PWD         string
}

func Run(in Inputs) {
	cdir := in.PWD

	if !fileutil.Exists(path.Join(cdir, projectFile)) {
		color.Red("seems that your current dir isn't a terraform project.")
		color.Red("you can create one running [rit scaffold generate terraform aws]")
		os.Exit(1)
	}

	// main.tf current

	mfile := path.Join(cdir, maintfFile)
	mb, _ := fileutil.ReadFile(mfile)
	mcfg, diags := hclwrite.ParseConfig(mb, mfile, hcl.InitialPos)
	if len(diags) != 0 {
		color.Red("unexpected diagnostics")
		for _, diag := range diags {
			color.Red(fmt.Sprintf("- %s", diag))
		}
		os.Exit(1)
	}

	// main.tk eks
	dir, _ := os.Getwd()
	efile := path.Join(dir, mainEKStfFile)
	eb, _ := fileutil.ReadFile(efile)
	ecfg, diags := hclwrite.ParseConfig(eb, efile, hcl.InitialPos)
	if len(diags) != 0 {
		color.Red("unexpected diagnostics")
		for _, diag := range diags {
			color.Red(fmt.Sprintf("- %s", diag))
		}
		os.Exit(1)
	}

	var reqblk *hclwrite.Block
	mbody := mcfg.Body()
	for _, block := range mbody.Blocks() {
		if block.Type() == "terraform" {
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
		if block.Type() == "terraform" {
			for _, tf := range block.Body().Blocks() {
				if tf.Type() == "required_providers" {
					for n, a := range tf.Body().Attributes() {
						var tokens hclwrite.Tokens
						tokens = a.BuildTokens(tokens)
						reqblk.Body().SetAttributeRaw(n, tokens[2:len(tokens) -1])
					}
				}
			}
		}
	}

	fileutil.WriteFile(mfile, mcfg.Bytes())

	// main.tf others
	mf, err := os.OpenFile(mfile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		color.Red(fmt.Sprintf("error openning %q, detail: %q", mfile, err))
		os.Exit(1)
	}
	defer mf.Close()
	_, err = mf.Write([]byte(tpl.Maintf))
	if err != nil {
		color.Red(fmt.Sprintf("error appending file %q, detail: %q", mfile, err))
		os.Exit(1)
	}

	// variables
	t := template.Must(template.New("Var").Parse(tpl.Variable))
	varf := path.Join(cdir, variableQA)
	vf, err := os.OpenFile(varf, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		color.Red(fmt.Sprintf("error openning %q, detail: %q", varf, err))
		os.Exit(1)
	}
	defer vf.Close()
	err = t.Execute(vf, in)
	if err != nil {
		color.Red(fmt.Sprintf("error writing %q, detail: %q", varf, err))
		os.Exit(1)
	}

	// dns_zone
	dzdir := path.Join(cdir, dnsZoneModule)
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

	// helm_deps
	helmdir := path.Join(cdir, hemlDeps)
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

	// iam_k8s
	iamdir := path.Join(cdir, iamK8SModule)
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

	fmt.Println()
	color.Green(fmt.Sprintln("eks module configured successfully."))
	color.Green(fmt.Sprintln("now, you can run [make plan] to check the terraform plan"))
}
