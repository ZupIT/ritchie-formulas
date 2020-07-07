package main

import (
	"os"

	"github.com/ZupIT/ritchie-formulas/scaffold/spring-starter/src/pkg/application"
)

func main() {
	loadInputs().Run()
}

func loadInputs() CommandHandler {
	return application.Inputs{
		Type:         os.Getenv("TYPE"),
		Language:     os.Getenv("LANGUAGE"),
		BootVersion:  os.Getenv("BOOT_VERSION"),
		GroupId:      os.Getenv("GROUP_ID"),
		ArtifactId:   os.Getenv("ARTIFACT_ID"),
		Description:  os.Getenv("DESCRIPTION"),
		Packaging:    os.Getenv("PACKAGING"),
		JavaVersion:  os.Getenv("JAVA_VERSION"),
		Dependencies: os.Getenv("DEPENDENCIES"),
	}
}

type CommandHandler interface {
	Run()
}
