package kafkautil

import (
	"sort"

	"github.com/Shopify/sarama"
	"github.com/manifoldco/promptui"
)

type KafkaVersions map[string]sarama.KafkaVersion

func PromptVersion() sarama.KafkaVersion {
	versions := kafkaVersions()
	kk := make([]string, 0, len(versions))

	for k := range versions {
		kk = append(kk, k)
	}

	sort.Strings(kk)
	prompt := promptui.Select{
		Items: kk,
		Templates: &promptui.SelectTemplates{
			Label: "Select Kafka version: ",
		},
	}
	_, result, _ := prompt.Run()

	return versions[result]
}

func kafkaVersions() KafkaVersions {
	vv := make(KafkaVersions)
	vv["V0.11.x"] = sarama.V0_11_0_2
	vv["V1.x.x"] = sarama.V1_1_1_0
	vv["V2.x.x"] = sarama.V2_4_0_0

	return vv
}
