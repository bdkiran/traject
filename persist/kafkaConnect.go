package persist

import (
	"os"

	"github.com/bdkiran/traject/utils"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

//InitilizeProducer creates a new producer that will be used to send messages to our kafka queue
func InitilizeProducer(bootstrapServer string) {
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": bootstrapServer})
	if err != nil {
		utils.DefaultLogger.Error.Println(err)
		os.Exit(1)
	}
	utils.DefaultLogger.Info.Println(producer.String())

}

//ProduceMessage sends a message to the kafka broker/topic
func ProduceMessage(message []byte) {
	utils.DefaultLogger.Info.Println("Produce message function called.")

	// p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost:9094"})
	// if err != nil {
	// 	panic(err)
	// }

	// defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					utils.DefaultLogger.Info.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					utils.DefaultLogger.Info.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "t1"
	producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, nil)

	// Wait for message deliveries before shutting down. Dont want to include this, since it will kill the producer
	//producer.Flush(15 * 1000)
}
