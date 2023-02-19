package kafkaservice;

import (
	Kafka "github.com/segmentio/kafka-go"
	Repository "xmservice.com/repository"
	Entity "xmservice.com/entity"
	JSON "encoding/json"
	"fmt"
	"context"
)

type KafkaRecieverService struct {

}

var kafkaService KafkaRecieverService;
var kafkaEventRepo Repository.KafkaEventRepository = &Repository.KafkaEventRepo{};

func (kaf *KafkaRecieverService) ConnectToKafka(kafkaURL string, topic string) *Kafka.Reader {
	fmt.Println("Kafka  ===  XMSERICES");
	brokers := []string{kafkaURL};
	return Kafka.NewReader(Kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  "kafkaEvents-group",
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	});
}

func (kaf *KafkaRecieverService) ConsumeKafkaTopicForMotherBoardService(){
	
	var kafkaURL string = "143.198.143.199:9092";
	var topic string = "xmservices-kafkaEvents";
	var kafkaReader *Kafka.Reader = kafkaService.ConnectToKafka(kafkaURL, topic);

	for {
        fmt.Println("Start  ..... KafkaEvent");
		var errorKaf error;
		ctx := context.Background();
		message, err := kafkaReader.FetchMessage(ctx);
		if err != nil {
			fmt.Println(err.Error());
			errorKaf = err; 
		}

		if message.Value != nil {
			if string(message.Key) == "kafkaEvents" {
				var kafkaEvent Entity.KafkaEvent;
				mErr := JSON.Unmarshal(message.Value, &kafkaEvent);
				if mErr != nil {
					errorKaf = mErr; 
				}
			
				data, errD := kafkaEventRepo.CreateKafkaEvent(kafkaEvent);
				if errD != nil {
					errorKaf = errD; 
				}
			}
		}

		if errorKaf == nil {
			if err := kafkaReader.CommitMessages(ctx, message); err != nil {
				fmt.Println("failed to commit messages kafkafEvents:", err.Error())
			}
		}
	}
}











