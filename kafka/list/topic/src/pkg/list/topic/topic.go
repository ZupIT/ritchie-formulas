package topic

import (
	"fmt"
	"log"
	"strings"

	"github.com/Shopify/sarama"

	"github.com/ZupIT/ritchie-formulas/kafka/src/pkg/kafkautil"
)

type Inputs struct {
	Urls string
}

func NewInputs(urls string) *Inputs {
	return &Inputs{Urls: urls}
}

func (i *Inputs) Run() {
	c := sarama.NewConfig()
	c.Version = kafkautil.PromptVersion()

	ca, err := sarama.NewClusterAdmin(strings.Split(i.Urls, ","), c)
	if err != nil {
		log.Println(err)
		return
	}

	tt, _ := ca.ListTopics()

	for k, _ := range tt {
		fmt.Println(k)
	}
}
