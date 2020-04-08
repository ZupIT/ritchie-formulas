package main

import (
	"log"
	"os"

	"github.com/ZupIT/ritchie-formulas/scaffold/spring-starter/src/pkg/application"
	"github.com/ZupIT/ritchie-formulas/scaffold/spring-starter/src/pkg/dependencies/list"
)

const (
	createCmd = "create"
	listCmd   = "list"
)

func main() {
	loadInputs().Run()
}

func loadInputs() CommandHandler {
	command := os.Getenv("COMMAND")
	switch command {
	case createCmd:
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
	case listCmd:
		return list.Inputs{
			Command: command,
		}
	default:
		log.Println("Command not found")
	}
	return nil
}

type CommandHandler interface {
	Run()
}
