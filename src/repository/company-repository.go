package repository

import (
	"gorm.io/gorm"
	Entity "xmservice.com/entity"
	Migration "xmservice.com/migration"
	"errors"
)

var dBs Migration.Migration = &Migration.MigrationService{};

type CompanyRepository interface {
	CreateCompany(company Entity.Company) (Entity.Company, error)
	UpdateCompany(company Entity.Company) (Entity.Company, error)
	GetCompanyId(id int32) (Entity.Company, error)
	GetCompanyByUuid(companyUuid string) (Entity.Company, error)
}

type CompanyRepo struct {

}

func (comp *CompanyRepo) CreateCompany(company Entity.Company) (Entity.Company, error){
	var database *gorm.DB = dBs.ConnectToDb();
	dbError := database.Create(&company).Error;
	if dbError != nil {
		return Entity.Company{}, errors.New(dbError.Error());
	}

	return company, nil;
}

func (comp *CompanyRepo) UpdateCompany(company Entity.Company) (Entity.Company, error){
	 var database *gorm.DB = dBs.ConnectToDb();
     dbError := database.Model(&Entity.Company{Id: company.Id}).Where(&Entity.Company{}).Updates(&company).Error;
	 if dbError != nil {
		 return Entity.Company{}, errors.New(dbError.Error());
	 }

	 return company, nil;
}

func (comp *CompanyRepo) GetCompanyId(id int32) (Entity.Company, error){
	var database *gorm.DB = dBs.ConnectToDb();
	var company Entity.Company;
	dbError := database.Where(&Entity.Company{Id: id}).Find(&company).Error;
	if dbError != nil {
		return Entity.Company{}, errors.New(dbError.Error());
	}

	return company, nil;
}

func (comp *CompanyRepo) GetCompanyByUuid(companyUuid string) (Entity.Company, error){
	var database *gorm.DB = dBs.ConnectToDb();
	var company Entity.Company;
	dbError := database.Where(&Entity.Company{CompanyUuid: companyUuid}).Find(&company).Error;
	if dbError != nil {
		return Entity.Company{}, errors.New(dbError.Error());
	}

	return company, nil;
}















