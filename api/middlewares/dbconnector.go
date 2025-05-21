package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/api/models"
	"github.com/words-reminder-api/internal/provider"
)

func DBConnector() gin.HandlerFunc {
	return func(c *gin.Context) {
		var container provider.ContainerRep = provider.GETContainer()
		var badResponse models.BadResponse

		db, err := container.MYSQLConnector().Open(c)
		if err != nil {
			badResponse.Code = http.StatusBadRequest
			badResponse.Error = "error openning db connection"
			c.IndentedJSON(http.StatusBadRequest, badResponse)
			return
		}

		fmt.Println("Success db connection!")
		c.Keys = map[string]any{"conn": db}

		c.Next()

		db.Close()

	}
}
