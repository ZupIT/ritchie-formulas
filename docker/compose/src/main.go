package main

import (
	"compose/pkg/compose"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
)

var questions = []*survey.Question{
	{
		Name: "letter",
		Prompt: &survey.MultiSelect{
			Message: "Select docker image: ",
			Options: []string{
				"awsclivl",
				"consul",
				"dynamoDB",
				"jaeger",
				"kafka",
				"mongo",
				"postgres",
				"redis",
				"rabbitmq",
			},
		},
	},
}

var qsPostgres = []*survey.Question{
	{
		Name: "dbname",
		Prompt: &survey.Input{
			Message: "Type DB name: ",
		},
		Validate:  survey.Required,
	},
	{
		Name: "dbuser",
		Prompt: &survey.Input{
			Message: "Type DB user: ",
		},
		Validate:  survey.Required,
	},
	{
		Name: "dbpass",
		Prompt: &survey.Password{
			Message: "Type DB password: ",
		},
		Validate:  survey.Required,
	},
}

var qsMongo = []*survey.Question{
	{
		Name: "mongoUser",
		Prompt: &survey.Input{
			Message: "Type Mongo WebClient user: ",
		},
		Validate:  survey.Required,
	},
	{
		Name: "mongoPass",
		Prompt: &survey.Password{
			Message: "Type Mongo WebClient password: ",
		},
		Validate:  survey.Required,
	},
}

var qsRabbitmq = []*survey.Question{
	{
		Name: "rabbitmqhost",
		Prompt: &survey.Input{
			Message: "Type Host name: ",
		},
		Validate:  survey.Required,
	},
	{
		Name: "rabbitmquser",
		Prompt: &survey.Input{
			Message: "Type RabbitMq user: ",
		},
		Validate:  survey.Required,
	},
	{
		Name: "rabbitmqpass",
		Prompt: &survey.Password{
			Message: "Type DB password: ",
		},
		Validate:  survey.Required,
	},
}

func main() {
	var selectItems []string
	extParams := make(map[string]string)

	err := survey.Ask(questions, &selectItems, survey.WithKeepFilter(true))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, item := range selectItems {
		switch item {
		case "postgres":
			answersPostgres := struct {
				DbName string
				DbUser string
				DbPass string
			}{}

			err := survey.Ask(qsPostgres, &answersPostgres)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			extParams["postgresDB"] = answersPostgres.DbName
			extParams["postgresUser"] = answersPostgres.DbUser
			extParams["postgresPassword"] = answersPostgres.DbPass

		case "mongo":
			answersMongo := struct {
				MongoUser string
				MongoPass string
			}{}

			err := survey.Ask(qsMongo, &answersMongo)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			extParams["mongoWebClientUser"] = answersMongo.MongoUser
			extParams["mongoWebClientPassword"] = answersMongo.MongoPass

		case "rabbitmq":
			answersRabitmq := struct {
				RabbitmqHost string
				RabbitmqUser string
				RabbitmqPass string
			}{}

			err := survey.Ask(qsRabbitmq, &answersRabitmq)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			extParams["rabbitmqHost"] = answersRabitmq.RabbitmqHost
			extParams["rabbitmqUser"] = answersRabitmq.RabbitmqUser
			extParams["rabbitmqPassword"] = answersRabitmq.RabbitmqPass
		}
	}

	compose.GenerateYml(selectItems, extParams)
}
