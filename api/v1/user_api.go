package api

import (
	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/controllers"
)

func UserEndpoints(r *gin.Engine) {
	r.GET("/albums", controllers.GetAlbums)
	r.POST("/albums", controllers.CreateNewUser)
}
