package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/api/controllers"
)

func SessionRoutes(rg *gin.RouterGroup) {
	sr := rg.Group("/session")
	sr.POST("/login/", controllers.Login)
}
