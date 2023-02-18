package api 

import (
	CompanyService "xmservice.com/services"
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

var companyService CompanyService.CompanyService;

type CompanyApi struct {

}

func (comA *CompanyApi) Router(router *gin.Engine){
	var route *gin.RouterGroup = router.Group("/api/company");
	var companyApi CompanyApi;
    companyApi.CreateCompany(route);
	companyApi.DeleteCompany(route);
	
}

func (comA *CompanyApi) CreateCompany(route *gin.RouterGroup){
	route.POST("/create", func(c *gin.Context) {

		data, err := companyService.CreateCompany(c);
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"status": http.StatusInternalServerError,
				"time": time.Now(),
			});

			return;
		}

		c.JSON(200, gin.H{
			"message": "Company Created Successfully",
			"data": data,
			"status": http.StatusOK,
		});
		
	});
}

func (comA *CompanyApi) PatchCompany(route *gin.RouterGroup){
	route.PATCH("/patch", func(c *gin.Context) {

		data, err := companyService.PatchCompany(c);
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"status": http.StatusInternalServerError,
				"time": time.Now(),
			});

			return;
		}

		c.JSON(200, gin.H{
			"message": "Company Updated Successfully",
			"data": data,
			"status": http.StatusOK,
		});
		
	});
}

func (comA *CompanyApi) DeleteCompany(route *gin.RouterGroup){
	route.DELETE("/delete", func(c *gin.Context) {

		_, err := companyService.PatchCompany(c);
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"status": http.StatusInternalServerError,
				"time": time.Now(),
			});

			return;
		}

		c.JSON(200, gin.H{
			"message": "Company Deleted Successfully",
			"data": "",
			"status": http.StatusOK,
		});
		
	});
}

func (comA *CompanyApi) GetCompany(route *gin.RouterGroup){
	route.GET("/id", func(c *gin.Context) {

		data, err := companyService.GetCompany(c);
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"status": http.StatusInternalServerError,
				"time": time.Now(),
			});

			return;
		}

		c.JSON(200, gin.H{
			"message": "",
			"data": data,
			"status": http.StatusOK,
		});
	});
}





