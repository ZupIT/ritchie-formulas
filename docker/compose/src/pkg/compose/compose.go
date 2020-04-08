package compose

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	headYml = `---
version: '2'
services:`

	kafkaYml = `  zookeeper:
    image: confluentinc/cp-zookeeper:5.4.0
    hostname: zookeeper
    container_name: zookeeper
    ports:
      - "2181:2181"
    environment:
      ZOOKEEPER_CLIENT_PORT: 2181
      ZOOKEEPER_TICK_TIME: 2000

  kafka:
    image: confluentinc/cp-server:5.4.0
    hostname: kafka
    container_name: kafka
    depends_on:
      - zookeeper
    ports:
      - "9092:9092"
    environment:
      KAFKA_kafka_ID: 1
      KAFKA_AUTO_CREATE_TOPICS_ENABLE: 'true'
      KAFKA_ZOOKEEPER_CONNECT: 'zookeeper:2181'
      KAFKA_LISTENER_SECURITY_PROTOCOL_MAP: PLAINTEXT:PLAINTEXT,PLAINTEXT_HOST:PLAINTEXT
      KAFKA_ADVERTISED_LISTENERS: PLAINTEXT://kafka:29092,PLAINTEXT_HOST://localhost:9092
      KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR: 1
      KAFKA_GROUP_INITIAL_REBALANCE_DELAY_MS: 0
      KAFKA_CONFLUENT_LICENSE_TOPIC_REPLICATION_FACTOR: 1`

	postgresYml = `  postgres:
    image: postgres:9.6
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_DB={{POSTGRES_DB}}
      - POSTGRES_USER={{POSTGRES_USER}}
      - POSTGRES_PASSWORD={{POSTGRES_PASSWORD}}
      - MAX_CONNECTIONS=300`

	mongoYml = `  mongo:
    image: mongo
    restart: always
    container_name: mongo
    volumes:
      - ./mongo_data:/data/db
    ports:
      - "27017:27017"

  mongo-express:
    image: mongo-express
    restart: always
    container_name: mongo-express
    ports:
      - "8081:8081"
    depends_on:
      - mongo
    links:
      - mongo:mongo
    environment:
      - ME_CONFIG_BASICAUTH_USERNAME={{MONGO_WEBCLIENT_USER}}
      - ME_CONFIG_BASICAUTH_PASSWORD={{MONGO_WEBCLIENT_PASSWORD}}
      - ME_CONFIG_MONGODB_SERVER=mongo
      - ME_CONFIG_MONGODB_PORT=27017
      - ME_CONFIG_MONGODB_ENABLE_ADMIN=true`

	stubby4jYml = `  stubby4j:
    image: joncanning/stubby4j
    ports:
      - "8787:8787"
      - "8882:8882"
    environment:
      STUBBY_PORT: 8882
    volumes:
      - ./files/stubby4j/integrations.yml:/usr/local/stubby.yml`

	jaegerYml = `  jaeger:
    image: jaegertracing/all-in-one:latest
    ports:
      - "5775:5775/udp"
      - "16686:16686"`

	dynamodbYml = `  dynamodb:
    image: amazon/dynamodb-local
    ports:
      - "8000:8000"`

	awscliv1Yml = `  awscliv1:
    image: normandesjr/awscliv1
    volumes:
      - ./files/dynamodb:/json
    command: create-resources.sh http://dynamodb:8000
    depends_on:
      - dynamodb`

	redisYml = `  redis:
    image: redis
    command: redis-server
    ports:
      - "6379:6379"`

	consulYml = `  consul:
    image: consul
    ports:
      - 8500:8500`

	integretionYml = `# Example
- request:
    method: GET
    url: ^/$

  response:
    status: 200
    
- request:
    method: POST
    url: ^/$

  response:
    status: 200`
)

func GenerateYml(items []string, extParams map[string]string) {
	ymlString := headYml + "\n"
	if len(items) == 1 {
		fmt.Println("No docker image selected")
	} else {
		for _, item := range items {
			switch item {
			case "kafka":
				ymlString = fmt.Sprintf("%s%s\n\n", ymlString, kafkaYml)
			case "postgres":
				postgresYmlString := postgresYml
				postgresYmlString = strings.Replace(postgresYmlString, "{{POSTGRES_DB}}", extParams["postgresDB"], -1)
				postgresYmlString = strings.Replace(postgresYmlString, "{{POSTGRES_USER}}", extParams["postgresUser"], -1)
				postgresYmlString = strings.Replace(postgresYmlString, "{{POSTGRES_PASSWORD}}", extParams["postgresPassword"], -1)
				ymlString = fmt.Sprintf("%s%s\n\n", ymlString, postgresYmlString)
			case "mongo":
				mongoYmlString := mongoYml
				mongoYmlString = strings.Replace(mongoYmlString, "{{MONGO_WEBCLIENT_USER}}", extParams["mongoWebClientUser"], -1)
				mongoYmlString = strings.Replace(mongoYmlString, "{{MONGO_WEBCLIENT_PASSWORD}}", extParams["mongoWebClientPassword"], -1)
				ymlString = fmt.Sprintf("%s%s\n\n", ymlString, mongoYmlString)
			case "stubby4j":
				createIfNotExists("files/stubby4j", 0755)
				writeFile("files/stubby4j/integrations.yml", []byte(integretionYml))
				ymlString = fmt.Sprintf("%s%s\n\n", ymlString, stubby4jYml)
			case "jaeger":
				ymlString = fmt.Sprintf("%s%s\n\n", ymlString, jaegerYml)
			case "dynamoDB":
				ymlString = fmt.Sprintf("%s%s\n\n", ymlString, dynamodbYml)
			case "awsclivl":
				ymlString = fmt.Sprintf("%s%s\n\n", ymlString, awscliv1Yml)
			case "redis":
				ymlString = fmt.Sprintf("%s%s\n\n", ymlString, redisYml)
			case "consul":
				ymlString = fmt.Sprintf("%s%s\n\n", ymlString, consulYml)
			}
		}
		writeFile("docker-compose.yml", []byte(ymlString))
		fmt.Println("Generated files in the current directory")
		fmt.Println("Run:\ndocker-compose up")
	}
}

// WriteFile wrapper for ioutil.WriteFile
func writeFile(path string, content []byte) error {
	return ioutil.WriteFile(path, content, 0644)
}

func createIfNotExists(dir string, perm os.FileMode) error {
	if exists(dir) {
		return nil
	}

	if err := os.MkdirAll(dir, perm); err != nil {
		return fmt.Errorf("failed to create directory: '%s', error: '%s'", dir, err.Error())
	}

	return nil
}

func exists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}

	return true
}
