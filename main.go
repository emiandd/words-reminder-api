package main

import (
	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/api/v1"
)

func main() {
	start()
}

func start() {
	// Specify connection properties.

	r := gin.Default()

	api.UserEndpoints(r)

	r.Run("localhost:8080")
}
