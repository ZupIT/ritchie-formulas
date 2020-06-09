package ritchie

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

type Input struct {
	KeycloakRelease string
	VaultRelease    string
	ConsulRelease   string
	RitchieRelease  string
	Namespace       string
	ValuesDir       string
}

func (in Input) Executor() {

	commands := in.getCommands()

	for _, cmd := range commands {
		log.Printf("%s", cmd.Args)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "APPLICATION=ritchie")
		cmd.Env = append(cmd.Env, "VAULT_CONTAINER="+in.VaultRelease+"-0")
		err := cmd.Run()
		if err != nil {
			log.Printf("%s", out.String())
			log.Fatal(err)
		}
		fmt.Printf("out: %q\n", out.String())
		time.Sleep(10 * time.Second)
	}

}

func (in Input) getCommands() []*exec.Cmd {
	return []*exec.Cmd{
		exec.Command("helm", "upgrade", "--install", in.ConsulRelease, "consul", "-f", "../consul/values.yaml", "--namespace", in.Namespace),
		exec.Command("../files/consulhealthy.sh"),
		exec.Command("helm", "upgrade", "--install", in.VaultRelease, "vault", "-f", in.ValuesDir+"/vault-values.yaml", "--namespace", in.Namespace),
		exec.Command("../files/configure.sh"),
		exec.Command("../files/unseal.sh"),
		exec.Command("kubectl", "-n", in.Namespace, "create", "secret", "generic", "realm-secret", "--from-file=realm-secret="+in.ValuesDir+"/ritchie-kc.json"),
		exec.Command("helm", "upgrade", "--install", in.KeycloakRelease, "keycloak", "-f", in.ValuesDir+"/keycloak-values.yaml", "--namespace", in.Namespace),
		exec.Command("helm", "upgrade", "--install", in.RitchieRelease, "ritchie-server", "-f", in.ValuesDir+"/ritchie-values.yaml", "--namespace", in.Namespace),
	}
}
