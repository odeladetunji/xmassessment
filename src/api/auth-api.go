package api 

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	Auth "xmservice.com/auth"
)

type AuthApi struct {

}

var authenticationService Auth.AuthenticationService;

func (autha *AuthApi) Router(router *gin.Engine){
	var route *gin.RouterGroup = router.Group("/api/auth");
	var authApi AuthApi;
    authApi.CreateToken(route);
}

func (autha *AuthApi) CreateToken(route *gin.RouterGroup){
	route.POST("/create", func(c *gin.Context) {

		data, err := authenticationService.CreateToken(c);
		if err != nil {
			c.JSON(500, gin.H{
				"message": err.Error(),
				"status": http.StatusInternalServerError,
				"time": time.Now(),
			});

			return;
		}

		c.JSON(200, gin.H{
			"message": "User token generated successfully!",
			"data": data,
			"status": http.StatusOK,
		});
		
	});
}



