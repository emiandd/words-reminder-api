package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/api/models"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var jwtKey = []byte("access_token")
		var badResponse models.BadResponse

		fmt.Println("Authenticating...")
		tokenString := c.Request.Header.Get("access_token")

		if tokenString == "" {
			badResponse.Code = http.StatusUnauthorized
			badResponse.Error = "empty token"
			c.AbortWithStatusJSON(badResponse.Code, badResponse)
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return jwtKey, nil
		})

		fmt.Println("token", token)

		if err != nil || !token.Valid {
			badResponse.Code = http.StatusUnauthorized
			badResponse.Error = "unauthorized"
			c.AbortWithStatusJSON(badResponse.Code, badResponse)
			return
		}

		c.Next()

	}
}
