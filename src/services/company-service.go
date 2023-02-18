package services


import (
	"github.com/gin-gonic/gin"
	Repository "xmservice.com/repository"
	Entity "xmservice.com/entity"
	"errors"
	"strings"
	"fmt"
	"strconv"
	"time"
)

type CompanyService struct {

}

var companyRepo Repository.CompanyRepository = &Repository.CompanyRepo{};

func (comps *CompanyService) CreateCompany(c *gin.Context) (Entity.Company, error) {

	var company Entity.Company;
	if err := c.BindJSON(&company); err != nil {
		return Entity.Company{}, errors.New(err.Error())
	}

	switch (company.Type) {
		case "Corporations":
			break;
		case "NonProfit":
			break;
		case "Cooperative": 
			break;
		case "Sole Proprietorship":
			break;
		default: 
			return Entity.Company{}, errors.New("type must be one of the following Corporations, NonProfit, Cooperative or Sole Proprietorship")
	}

	if len(company.CompanyUuid) < 10 {
		return Entity.Company{}, errors.New("companyUuid is required, it has to be a string, minimum of 10 characters")
	}

	if len(company.CompanyName) != 15 {
		return Entity.Company{}, errors.New("companyName is required, it has to be a string of 15 characters")
	}

	if len(company.Description) > 3000 {
		return Entity.Company{}, errors.New("description cannot be greater than 3000 characters")
	}

	comp, errCC := companyRepo.GetCompanyByUuid(company.CompanyUuid);
	if errCC != nil {
		return Entity.Company{}, errors.New(errCC.Error());
	}

	if comp.Id != 0 {
		return Entity.Company{}, errors.New(strings.Join([]string{"company with uuid ", fmt.Sprint(comp.CompanyUuid), " already exits"}, ""));
	}

	company.CreatedDate = time.Now();
	company.CreatedBy = "Admin";
	company.LastActivityBy = "Admin";
	company.LastActivityDate = time.Now();

	company, errC := companyRepo.CreateCompany(company);
    if errC != nil {
		return Entity.Company{}, errors.New(errC.Error());
	}

	return company, nil;

}

func (comps *CompanyService) PatchCompany(c *gin.Context) (Entity.Company, error) {

	var acompany Entity.Company;
	if err := c.BindJSON(&acompany); err != nil {
		return Entity.Company{}, errors.New(err.Error())
	}

	switch (acompany.Type) {
		case "Corporations":
			break;
		case "NonProfit":
			break;
		case "Cooperative": 
			break;
		case "Sole Proprietorship":
			break;
		default: 
			return Entity.Company{}, errors.New("type must be one of the following Corporations, NonProfit, Cooperative or Sole Proprietorship")
	}

	if len(acompany.CompanyUuid) < 10 {
		return Entity.Company{}, errors.New("companyUuid is required, it has to be a string, minimum of 10 characters")
	}

	if len(acompany.CompanyName) != 15 {
		return Entity.Company{}, errors.New("companyName is required, it has to be a string of 15 characters")
	}

	if len(acompany.Description) > 3000 {
		return Entity.Company{}, errors.New("description cannot be greater than 3000 characters")
	}

	company, errC := companyRepo.GetCompanyId(acompany.Id);
    if errC != nil {
		return Entity.Company{}, errors.New(errC.Error());
	}

	if company.Id == 0 {
		return Entity.Company{}, errors.New(strings.Join([]string{"company with id ", fmt.Sprint(acompany.Id), "does not exits"}, ""));
	}

	company.Description = acompany.Description;
	company.CompanyName = acompany.CompanyName;
	company.NumberOfEmployees = acompany.NumberOfEmployees;
	company.Registered = acompany.Registered;
	company.Type = acompany.Type;
	company.LastActivityBy = "Admin";
	company.LastActivityDate = time.Now();

	company, errI := companyRepo.UpdateCompany(company);
    if errI != nil {
		return Entity.Company{}, errors.New(errI.Error());
	}

	return company, nil;

}

func (comps *CompanyService) DeleteCompany(c *gin.Context) (error) {

	var acompany Entity.Company;
	if err := c.BindJSON(&acompany); err != nil {
		return errors.New(err.Error())
	}

	company, err := companyRepo.GetCompanyId(acompany.Id);
    if err != nil {
		return errors.New(err.Error());
	}

	if company.Id == 0 {
		return errors.New(strings.Join([]string{"company with id ", fmt.Sprint(acompany.Id), "does not exits"}, ""));
	}

	company.IsDeleted = true;
	company.LastActivityBy = "Admin";
	company.LastActivityDate = time.Now();

	company, err = companyRepo.UpdateCompany(company);
    if err != nil {
		return errors.New(err.Error());
	}

	return nil;

}

func (comps *CompanyService) GetCompany(c *gin.Context) (Entity.Company, error) {
    
	if len(c.Query("id")) == 0 {
		return Entity.Company{}, errors.New("Id is required");
	}

	id, errU := strconv.Atoi(c.Query("id"));
	if errU != nil {
		return Entity.Company{}, errors.New(errU.Error());
	}

	company, err := companyRepo.GetCompanyId(int32(id));
    if err != nil {
		return Entity.Company{}, errors.New(err.Error());
	}

	if company.Id == 0 {
		return Entity.Company{}, errors.New(strings.Join([]string{"company with id ", fmt.Sprint(company.Id), "does not exits"}, ""));
	}

	return company, nil;

}

