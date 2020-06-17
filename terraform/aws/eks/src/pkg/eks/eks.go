package eks

import (
	"fmt"
	"html/template"
	"os"
	"path"

	"eks/pkg/tpl"

	"github.com/ZupIT/ritchie-cli/pkg/file/fileutil"
	"github.com/fatih/color"
	"github.com/hashicorp/terraform/configs"
)

const (
	projectFile   = ".scaffold"
	maintfFile    = "src/main.tf"
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

	parser := configs.NewParser(nil)
	cfg, diags := parser.LoadConfigFile(path.Join(cdir, maintfFile))
	if len(diags) != 0 {
		color.Red("unexpected diagnostics")
		for _, diag := range diags {
			color.Red(fmt.Sprintf("- %s", diag))
		}
		os.Exit(1)
	}

	exists := false
	for _, m := range cfg.ModuleCalls {
		if m.Name == "kubernetes_cluster" {
			exists = true
			break
		}
	}

	fmt.Println("main.tf:", cfg)

	if !exists {
		// main.tf
		mf, err := os.OpenFile(path.Join(cdir, maintfFile), os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			color.Red(fmt.Sprintf("error openning main.tf, detail: %q", err))
			os.Exit(1)
		}
		defer mf.Close()
		if _, err := mf.WriteString(tpl.Maintf); err != nil {
			color.Red(fmt.Sprintf("error writing main.tf, detail: %q", err))
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

}
