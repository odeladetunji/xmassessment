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

	company, errC := companyRepo.GetCompany(acompany.Id);
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

	company, err := companyRepo.GetCompany(acompany.Id);
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

	company, err := companyRepo.GetCompany(int32(id));
    if err != nil {
		return Entity.Company{}, errors.New(err.Error());
	}

	if company.Id == 0 {
		return Entity.Company{}, errors.New(strings.Join([]string{"company with id ", fmt.Sprint(company.Id), "does not exits"}, ""));
	}

	return company, nil;

}

