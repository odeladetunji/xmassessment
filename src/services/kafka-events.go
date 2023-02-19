package services


import (
	"github.com/gin-gonic/gin"
	Repository "xmservice.com/repository"
	Entity "xmservice.com/entity"
	"errors"
	"time"
)

type KafkaEventService struct {

}

var kafkaEventRepo Repository.KafkaEventRepository = &Repository.KafkaEventRepo{};

func (kafkas *KafkaEventService) GetAllKafkaEvents(c *gin.Context) ([]Entity.KafkaEvent, error) {
	
	eventList, errE := kafkaEventRepo.GetAllKafkaEvents();
	if errE != nil {
		return []Entity.KafkaEvent, errors.New(errE.Error());
	}

	return eventList, nil;

}

