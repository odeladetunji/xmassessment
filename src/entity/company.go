package entity

import (
	"time"
	// "gopkg.in/go-playground/validator.v9"
)

type Company struct {
	Id int32 `json:"id"`
	CompanyUuid string `json:"companyUuid" validate:"required"`
	CompanyName string `json:"companyName" validate:"required"`
	Description string `json:"description"`
	IsDeleted bool `json:"isDeleted"`
	NumberOfEmployees int32 `json:"numberOfEmployees" validate:"required"`
	Registered bool `json:"registered" validate:"required"`
	Type string `json:"type" validate:"required"`
	CreatedDate time.Time `json:"createdDate"`
	CreatedBy string `json:"createdBy"`
	LastActivityBy string `json:"lastActivityBy"`
	LastActivityDate time.Time `json:"lastActivityDate"`
}






