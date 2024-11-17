package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/words-reminder-api/api/middlewares"
	"github.com/words-reminder-api/api/routes"
)

func main() {
	start()
}

func start() {
	_ = godotenv.Load(".env")
	r := gin.Default()

	// Agrega middleware CORS
	config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Access-Control-Allow-Origin", "Origin", "X-Requested-With", "Content-Type", "Accept"}
	r.Use(cors.New(config))
	r.Use(middlewares.Authenticate())

	// r.Use(middlewares.Cors())
	r.Use(middlewares.DBConnector())
	r = SetupRouter(r)
	getRoutes(r)
	r.Run("localhost:8080")
}

func SetupRouter(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func getRoutes(r *gin.Engine) {
	v1 := r.Group("/v1")
	routes.UserRoutes(v1)
	routes.SessionRoutes(v1)
	routes.WordRoutes(v1)

	// addPingRoutes(v1)

	// v2 := router.Group("/v2")
	// addPingRoutes(v2)
}
