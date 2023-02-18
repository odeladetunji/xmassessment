package entity

import "time"

type Company struct {
	Id int32 `json:"id"`
	CompanyName string `json:"companyName"`
	Description string `json:"description"`
	IsDeleted bool `json:"isDeleted"`
	NumberOfEmployees int32 `json:"numberOfEmployees"`
	Registered bool `json:"registered"`
	Type string `json:"type"`
	CreatedDate time.Time `json:"createdDate"`
	CreatedBy string `json:"createdBy"`
	LastActivityBy string `json:"lastActivityBy"`
	LastActivityDate time.Time `json:"lastActivityDate"`
}
