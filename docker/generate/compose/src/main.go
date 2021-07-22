package main

import (
	"compose/pkg/compose"
	"compose/pkg/prompt"
)

func main() {
	var selectItems []string
	selectItem := ""
	var extParams = make(map[string]string)
	items := []string{"awsclivl", "consul", "dynamoDB", "jaeger", "kafka", "mongo", "postgres", "mysql", "redis", "stubby4j", "rabbitmq", "finish!"}
	for selectItem != "finish!" {
		selectItem, _ = prompt.List("Select docker image: ", items)
		if selectItem == "postgres" {
			extParams["postgresDB"], _ = prompt.String("Type DB name: ", true)
			extParams["postgresUser"], _ = prompt.String("Type DB user: ", true)
			extParams["postgresPassword"], _ = prompt.StringPwd("Type DB password: ")
		}
		if selectItem == "mysql" {
			extParams["mysqlDB"], _ = prompt.String("Type DB name: ", true)
			extParams["mysqlUser"], _ = prompt.String("Type DB user: ", true)
			extParams["mysqlPassword"], _ = prompt.StringPwd("Type DB password: ")
		}
		if selectItem == "mongo" {
			extParams["mongoWebClientUser"], _ = prompt.String("Type Mongo WebClient user: ", true)
			extParams["mongoWebClientPassword"], _ = prompt.StringPwd("Type Mongo WebClient password: ")
		}
		if selectItem == "rabbitmq" {
			extParams["rabbitmqHost"], _ = prompt.String("Type Host name: ", true)
			extParams["rabbitmqUser"], _ = prompt.String("Type RabbitMq user: ", true)
			extParams["rabbitmqPassword"], _ = prompt.StringPwd("Type RabbitMq password: ")
		}
		selectItems = append(selectItems, selectItem)
		for i, item := range items {
			if item == selectItem { //Remove input to list
				items = append(items[:i], items[i+1:]...)
				break
			}
		}
	}

	compose.GenerateYml(selectItems, extParams)
}
