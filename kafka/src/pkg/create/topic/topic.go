package topic

import (
	"log"
	"strings"

	"github.com/Shopify/sarama"

	"github.com/ZupIT/ritchie-formulas/kafka/src/pkg/kafkautil"
)

type Inputs struct {
	Urls        string
	Name        string
	Replication int16
	Partitions  int32
}

func NewInputs(urls, name string, replication int16, partitions int32) *Inputs {
	return &Inputs{Urls: urls, Name: name, Replication: replication, Partitions: partitions}
}

func (i *Inputs) Run() {
	c := sarama.NewConfig()
	c.Version = kafkautil.PromptVersion()

	ca, err := sarama.NewClusterAdmin(strings.Split(i.Urls, ","), c)
	if err != nil {
		log.Println(err)
		return
	}

	d := sarama.TopicDetail{NumPartitions: i.Partitions, ReplicationFactor: i.Replication}
	err = ca.CreateTopic(i.Name, &d, false)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Kafka topic created successfully!")
}
