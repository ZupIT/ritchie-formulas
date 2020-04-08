package produce

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Shopify/sarama"

	"github.com/ZupIT/ritchie-formulas/kafka/src/pkg/kafkautil"
)

const delimiter = '\n'

type Inputs struct {
	Urls  string
	Topic string
}

func NewInputs(urls, topic string) *Inputs {
	return &Inputs{Urls: urls, Topic: topic}
}

func (i *Inputs) Run() {
	c := sarama.NewConfig()
	c.Version = kafkautil.PromptVersion()
	c.Producer.Return.Successes = true
	c.Producer.RequiredAcks = sarama.WaitForAll
	c.Producer.Retry.Max = 5

	p, err := sarama.NewSyncProducer(strings.Split(i.Urls, ","), c)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := p.Close(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Headers (e.g., key=value,key=value): ")
		h, _ := r.ReadString(delimiter)

		fmt.Print("Key: ")
		k, _ := r.ReadString(delimiter)

		fmt.Print("Message: ")
		m, _ := r.ReadString(delimiter)

		publish(p, i.Topic, removeSuffix(h), removeSuffix(k), removeSuffix(m))
	}

}

func publish(producer sarama.SyncProducer, topic string, headers, key, msg string) {
	m := &sarama.ProducerMessage{
		Topic: topic,
	}

	if headers != "" {
		hh, err := parseHeader(headers)
		if err != nil {
			fmt.Println(err)
			return
		}
		m.Headers = hh
	}

	if key != "" {
		m.Key = sarama.StringEncoder(key)
	}

	if msg == "" {
		fmt.Println("The field message not be empty!")
		return
	}
	m.Value = sarama.StringEncoder(msg)

	_, _, err := producer.SendMessage(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("The message was successfully stored!\n")
}

func parseHeader(headers string) ([]sarama.RecordHeader, error) {
	hh := strings.Split(headers, ",")

	rr := make([]sarama.RecordHeader, len(hh))

	for i, s := range hh {
		kv := strings.Split(s, "=")
		if len(kv) < 2 {
			return nil, errors.New("invalid kafka headers, use this format (key=value)")
		}

		rr[i] = sarama.RecordHeader{Key: []byte(kv[0]), Value: []byte(kv[1])}
	}

	return rr, nil
}

func removeSuffix(text string) string {
	return strings.TrimSuffix(text, string(delimiter))
}
