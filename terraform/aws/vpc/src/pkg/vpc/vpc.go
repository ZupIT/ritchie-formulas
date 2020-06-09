package vpc

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/fatih/color"
)

const (
	vpcAZSFormat = "\"%s\"%s"
)

var ()

type Inputs struct {
	Region       string
	VPCName      string
	VPCCIDR      string
	VPCAZS       string
	CustomerName string
}

func Run(in Inputs) {

	fmt.Println()
	color.HiBlack(fmt.Sprintln("vpc/main.tf"))
	color.HiBlack(fmt.Sprintln("============================================================="))
	fmt.Println(Main)

	ss := strings.Split(in.VPCAZS, ",")
	siz := len(ss) - 1
	var azs string
	for i, s := range ss {
		if i < siz {
			azs += fmt.Sprintf(vpcAZSFormat, s, ",")
		} else {
			azs += fmt.Sprintf(vpcAZSFormat, s, "")
		}
	}
	in.VPCAZS = azs

	fmt.Println()
	color.HiBlack(fmt.Sprintln("vpc/variables.tf"))
	color.HiBlack(fmt.Sprintln("============================================================="))
	t := template.Must(template.New("Vars").Parse(Vars))
	err := t.Execute(color.Output, in)
	if err != nil {
		fmt.Errorf("executing template:\n", err)
	}
	fmt.Println()
	color.Green(fmt.Sprintln("tf files generated successfully."))
	color.Green(fmt.Sprintln("Localtion: ./build/2593f872-f430-463e-858a-34d2ef30c5bc"))
	color.Green(fmt.Sprintln("Now, you can run [terraform plan|apply]"))
}
