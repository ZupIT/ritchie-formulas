package generate

import (
	"fmt"
	"generate/pkg/fileutil"
	"generate/pkg/generate/tpl"
	"os"
	"strings"
)

type Inputs struct {
	VaultRelease            string
	ConsulRelease           string
	KeycloakExternalAddress string
	KeycloakAdminPassword   string
	KeycloakRealm           string

	KeycloakClientId     string
	KeycloakClientSecret string
	KeycloakDbUser       string
	KeycloakDbPassword   string
	KeycloakDbAddress    string

	KeycloakDbDatabase        string
	RitchieCliVersionProvider string
	RitchieCliVersionUrl      string
	RitchieOrganization       string
	RitchieRepositories       string
}

func (in Inputs) Files() {
	fmt.Println("Generating Initial Files...")
	dirTpl := "templates"

	fileutil.CreateDirIfNotExists(dirTpl, os.ModePerm)

	tplRitchieFile := tpl.TplRitchieJson

	tplRitchieFile = strings.ReplaceAll(tplRitchieFile, "{{ keycloakRealm }}", in.KeycloakRealm)
	tplRitchieFile = strings.ReplaceAll(tplRitchieFile, "{{ keycloakClientId }}", in.KeycloakClientId)
	tplRitchieFile = strings.ReplaceAll(tplRitchieFile, "{{ keycloakClientSecret }}", in.KeycloakClientSecret)

	fileutil.WriteFile(fmt.Sprintf("%s/ritchie-kc.json", dirTpl), []byte(tplRitchieFile))

	tplKeyCloak := tpl.TplKeycloakValues
	tplKeyCloak = strings.ReplaceAll(tplKeyCloak, "{{ keycloakDbUser }}", in.KeycloakDbUser)
	tplKeyCloak = strings.ReplaceAll(tplKeyCloak, "{{ keycloakDbPassword }}", in.KeycloakDbPassword)
	tplKeyCloak = strings.ReplaceAll(tplKeyCloak, "{{ keycloakDbAddress }}", in.KeycloakDbAddress)
	tplKeyCloak = strings.ReplaceAll(tplKeyCloak, "{{ keycloakDbDatabase }}", in.KeycloakDbDatabase)
	tplKeyCloak = strings.ReplaceAll(tplKeyCloak, "{{ keycloakAdminPassword }}", in.KeycloakAdminPassword)

	fileutil.WriteFile(fmt.Sprintf("%s/keycloak-values.yaml", dirTpl), []byte(tplKeyCloak))

	tplRitchieValues := tpl.TplRitchieValues
	//TODO : Add Tag
	tplRitchieValues = strings.ReplaceAll(tplRitchieValues, "{{ vaultRelease }}", in.VaultRelease)
	tplRitchieValues = strings.ReplaceAll(tplRitchieValues, "{{ keycloakRealm }}", in.KeycloakRealm)
	tplRitchieValues = strings.ReplaceAll(tplRitchieValues, "{{ keycloakClientId }}", in.KeycloakClientId)
	tplRitchieValues = strings.ReplaceAll(tplRitchieValues, "{{ keycloakClientSecret }}", in.KeycloakClientSecret)
	tplRitchieValues = strings.ReplaceAll(tplRitchieValues, "{{ ritchieOrganization }}", in.RitchieOrganization)
	tplRitchieValues = strings.ReplaceAll(tplRitchieValues, "{{ ritchieCliVersionProvider }}", in.RitchieCliVersionProvider)
	tplRitchieValues = strings.ReplaceAll(tplRitchieValues, "{{ ritchieCliVersionURL }}", in.RitchieCliVersionUrl)
	tplRitchieValues = strings.ReplaceAll(tplRitchieValues, "{{ ritchieRepositories }}", in.RitchieRepositories)
	tplRitchieValues = strings.ReplaceAll(tplRitchieValues, "{{ keycloakExternalAddress }}", in.KeycloakExternalAddress)

	fileutil.WriteFile(fmt.Sprintf("%s/ritchie-values.yaml", dirTpl), []byte(tplRitchieValues))

	tplVaultValues := tpl.TplVaultValues

	tplVaultValues = strings.ReplaceAll(tplVaultValues, "{{ consulRelease }}", in.ConsulRelease)
	fileutil.WriteFile(fmt.Sprintf("%s/vault-values.yaml", dirTpl), []byte(tplVaultValues))

	fmt.Println("Finish ...")
}
