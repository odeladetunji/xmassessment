package repository

import (
	"gorm.io/gorm"
	Entity "xmservice.com/entity"
	Migration "xmservice.com/migration"
	"errors"
)

var dtBK Migration.Migration = &Migration.MigrationService{};

type KafkaEventRepository interface {
	CreateKafkaEvent(kafkaEvent Entity.KafkaEvent) error 
	GetAllKafkaEvents() ([]Entity.KafkaEvent, error)
}

type KafkaEventRepo struct {
    
}

func (kafk *KafkaEventRepo) CreateKafkaEvent(kafkaEvent Entity.KafkaEvent) error {
	var database *gorm.DB = dtBK.ConnectToDb();
	dbError := database.Create(&kafkaEvent).Error;
	if dbError != nil {
		return errors.New(dbError.Error());
	}

	return nil;
}

func (kafk *KafkaEventRepo) GetAllKafkaEvents() ([]Entity.KafkaEvent, error){
	var database *gorm.DB = dtBK.ConnectToDb();
	var kafkaEventList []Entity.KafkaEvent;
	dbError := database.Find(&kafkaEventList).Error;
	if dbError != nil {
		return []Entity.KafkaEvent{}, errors.New(dbError.Error());
	}

	return kafkaEventList, nil;
}


















