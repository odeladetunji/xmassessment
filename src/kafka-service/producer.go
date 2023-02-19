package kafkaservice;

import (
	Kafka "github.com/segmentio/kafka-go"
	Entity "xmservice.com/entity"
	"github.com/joho/godotenv"
	"context"
	"errors"
	"strconv"
	"os"
	"log"
	"net"
	JSON "encoding/json"
)

type KafkaService struct {

}

func (kaf *KafkaService) CreateKafkaConnection(kafkaUrl string, topic string) (*Kafka.Writer) {
	return &Kafka.Writer{
		Addr:     Kafka.TCP(kafkaUrl),
		Topic:    topic,
		Balancer: &Kafka.LeastBytes{},
	}
}

func (kaf *KafkaService) KafkaProducerCreateKafkaEvent(kafkaWriter *Kafka.Writer, key string, byteArray []byte) {

	message := Kafka.Message{
		Key: []byte(key),
		Value: []byte(byteArray),
	}

	// var req *http.Request;
	errI := kafkaWriter.WriteMessages(context.Background(), message);
	if errI != nil {
		panic(errI.Error());
	}

}

func (kaf *KafkaService) PushToKafkaProducer(kafkaEvent *Entity.KafkaEvent, topicType string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// var kafkaUrl string = "143.198.143.199:9092";
	var kafkaURL string = os.Getenv("KAFKA_SERVER");
	var topic string = "xmservices-kafkaEvents";
	var key string;
	var byteArray []byte;

	if topicType == "CREATE-KAFKA-EVENTS" {
		key = "kafkaEvents";
		json, err := JSON.Marshal(kafkaEvent);
		if err != nil {
			panic(err.Error());
		}

		byteArray = json;
	}

	//Create Connection
	var kafkaService KafkaService;
	kafkaWriter := kafkaService.CreateKafkaConnection(kafkaURL, topic);
	defer kafkaWriter.Close();

	//Push to Producer
	kafkaService.KafkaProducerCreateKafkaEvent(kafkaWriter, key, byteArray);

}

func (kaf *KafkaService) checkIfTopicExists(topic string) (bool, error){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// var kafkaUrl string = "143.198.143.199:9092";
	var kafkaURL string = os.Getenv("KAFKA_SERVER")
	conn, err := Kafka.Dial("tcp", kafkaURL);
	if err != nil {
		return false, errors.New(err.Error());
	}
	defer conn.Close()

	partitions, errR := conn.ReadPartitions()
	if errR != nil {
		return false, errors.New(errR.Error());
	}

	m := map[string]struct{}{}

	for _, p := range partitions {
		m[p.Topic] = struct{}{}
	}

	var topicIsPresent bool = false;
	for k := range m {
		if k == topic {
			topicIsPresent = true;
		}
	}

	return topicIsPresent, nil;

}

func (kaf *KafkaService) CreateKafKaTopic(){
	var kafkaService KafkaService;

	topicCreation := func(topic string){
		topicIsPresent, errTop := kafkaService.checkIfTopicExists(topic);
		if errTop != nil {
			panic(errTop.Error())
		}

		if errTop == nil && !topicIsPresent {
			err := godotenv.Load()
			if err != nil {
				log.Fatal("Error loading .env file")
			}

	       // var kafkaUrl string = "143.198.143.199:9092";
		    var kafkaURL string = os.Getenv("KAFKA_SERVER")
			conn, err := Kafka.Dial("tcp", kafkaURL)
			if err != nil {
				panic(err.Error())
			}
			defer conn.Close();
		
			controller, err := conn.Controller()
			if err != nil {
				panic(err.Error())
			}
			var controllerConn *Kafka.Conn
			controllerConn, err = Kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
			if err != nil {
				panic(err.Error())
			}
	
			defer controllerConn.Close();
		
			topicConfigs := []Kafka.TopicConfig{
				{
					Topic:             topic,
					NumPartitions:     1,
					ReplicationFactor: 1,
				},
			}
		
			err = controllerConn.CreateTopics(topicConfigs...);
			if err != nil {
				panic(err.Error())
			}
		}
	}

	var topicList []string = []string{"xmservices-kafkaEvents"};
    for i := 0; i < len(topicList); i++ {
		topicCreation(topicList[i]);
	}
}




































