package entity;

import "time"

type KafkaEvent struct {
	Id int32 `json:"id"`
	Data string `json:"data"`
	Type string `json:"type"`
	CreatedDate time.Time `json:"createdDate"`
	CreatedBy string `json:"createdBy"`
	LastActivityBy string `json:"lastActivityBy"`
	LastActivityDate time.Time `json:"lastActivityDate"`
}

