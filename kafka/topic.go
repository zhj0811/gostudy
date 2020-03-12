package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

func Create(topic *TopicInfo) error {
	log.Println("start create topic...")

	broker := sarama.NewBroker("localhost:9092")
	err := broker.Open(nil)
	if err != nil {
		log.Println("error in open the broker, ", err)
		return err
	}

	var topicDetail sarama.TopicDetail
	config := sarama.NewConfig()
	// config.Version = sarama.V2_1_0_0 //kafka版本号
	config.Net.SASL.Enable = true
	// config.Net.SASL.Mechanism = "PLAIN"
	config.Net.SASL.User = "admin"
	config.Net.SASL.Password = "admin"

	admin, err := sarama.NewClusterAdmin([]string{"192.168.10.128:9092"}, config)
	if err != nil {
		log.Fatal("error in create new cluster ... ", err)
		return err
	}

	err = admin.CreateTopic(topic.TopicName, &topicDetail, false)
	if err != nil {
		log.Println("error in create topic, ", err)
		return err
	}

	err = admin.Close()
	if err != nil {
		log.Fatal("error in close admin, ", err)
		return err
	}

	return nil
}

type TopicInfo struct {
	TopicName string
}

var name string

func init() {
	flag.StringVar(&name, "n", "test", "topic name")
}

func main() {
	flag.Parse()
	fmt.Println("topic name: ", name)
	topic := &TopicInfo{name}
	if err := Create(topic); err != nil {
		fmt.Println("create topic error: ", err.Error())
		return
	}
	fmt.Println("create topic success.")
	return
}
