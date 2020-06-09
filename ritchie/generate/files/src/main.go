package main

import (
	"generate/pkg/generate"
	"os"
)

func main() {
	generate.Inputs{
		VaultRelease:            os.Getenv("VAULT_RELEASE"),
		ConsulRelease:           os.Getenv("CONSUL_RELEASE"),
		KeycloakExternalAddress: os.Getenv("KEYCLOAK_EXTERNAL_ADDRESS"),
		KeycloakAdminPassword:   os.Getenv("KEYCLOAK_ADMIN_PASSWORD"),
		KeycloakRealm:           os.Getenv("KEYCLOAK_REALM"),

		KeycloakClientId:     os.Getenv("KEYCLOAK_CLIENT_ID"),
		KeycloakClientSecret: os.Getenv("KEYCLOAK_CLIENT_SECRET"),
		KeycloakDbUser:       os.Getenv("KEYCLOAK_DB_USER"),
		KeycloakDbPassword:   os.Getenv("KEYCLOAK_DB_PASSWORD"),
		KeycloakDbAddress:    os.Getenv("KEYCLOAK_DB_ADDRESS"),

		KeycloakDbDatabase:        os.Getenv("KEYCLOAK_DB_DATABASE"),
		RitchieCliVersionProvider: os.Getenv("RITCHIE_CLI_VERSION_PROVIDER"),
		RitchieCliVersionUrl:      os.Getenv("RITCHIE_CLI_VERSION_URL"),
		RitchieOrganization:       os.Getenv("RITCHIE_ORGANIZATION"),
		RitchieRepositories:       os.Getenv("RITCHIE_REPOSITORIES"),
	}.Files()

}
