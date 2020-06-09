package main

import (
	"os"
	"ritchie/pkg/ritchie"
)

func main() {

	ritchie.Input{
		KeycloakRelease: os.Getenv("KEYCLOAK_RELEASE"),
		ConsulRelease:   os.Getenv("CONSUL_RELEASE"),
		VaultRelease:    os.Getenv("VAULT_RELEASE"),
		RitchieRelease:  os.Getenv("RITCHIE_RELEASE"),
		Namespace:       os.Getenv("NAMESPACE"),
		ValuesDir:       os.Getenv("VALUES_DIR"),
	}.Executor()

}
