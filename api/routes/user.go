package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/api/controllers"
	m "github.com/words-reminder-api/api/middlewares"
)

func UserRoutes(rg *gin.RouterGroup) {
	ur := rg.Group("/users")
	ur.POST("/", controllers.CreateNewUser)
	ur.Use(m.Authenticate())
	ur.GET("/", controllers.GetUsers)
}
