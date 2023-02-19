
package auth;

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/vladimiroff/jwt-go/v3"
	"errors"
	"net/http"
	"time"
	// "fmt"
	"strings"
)

type AuthenticationService struct {

}

var mySignedKey = []byte("xmauthentication");

func (authen *AuthenticationService) CreateToken(c *gin.Context) (string, error) {

	type User struct {
		UserName string `json:"username"`
		Password string `json:"password"`
	}

	var user User;
	if err := c.BindJSON(&user); err != nil {
		return "", errors.New(err.Error())
	}

	if len(user.UserName) == 0 {
		return "", errors.New("username is required")
	}

	if len(user.Password) == 0 {
		return "", errors.New("password is required")
	}

	token := jwt.New(jwt.SigningMethodHS256);
	claims := token.Claims.(jwt.MapClaims);
    
	claims["authorized"] = true;
	claims["user"] = user.UserName;
	claims["password"] = user.Password;
	claims["role"] = "admin";
	claims["exp"] = time.Now().AddDate(0, 0, 1);

	tokenString, err := token.SignedString(mySignedKey);

	if err != nil {
		return "", errors.New(err.Error())
	}
    
	return tokenString, nil

}

func (authen *AuthenticationService) ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		if strings.Contains(c.Request.URL.Path, "/api/company/id") {
			c.Next();
			return;
		}

		var auth string = c.Request.Header.Get("Authorization");
		if len(auth) == 0 {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort();
			return
	    }

		var token string = strings.Split(auth, " ")[1];

		if len(token) == 0 {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort();
			return
		}

		_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return mySignedKey, nil
		});

		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		c.Next();
	}
}




