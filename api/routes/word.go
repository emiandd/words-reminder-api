package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/api/controllers"
	"github.com/words-reminder-api/api/middlewares"
)

func WordRoutes(rg *gin.RouterGroup) {
	ur := rg.Group("/words")
	ur.Use(middlewares.Authenticate())
	ur.POST("/", controllers.CreateNewWord)
	ur.GET("/", controllers.FetchWords)
}
