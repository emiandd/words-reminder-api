package middlewares

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Permitir solicitudes desde cualquier origen
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		// Permitir los métodos GET, POST, PUT, DELETE, OPTIONS
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// Permitir los encabezados Content-Type y Authorization
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Permitir que los navegadores envíen las cookies
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// // Si la solicitud es una prefligth (OPTIONS), finalizar aquí
		// if c.Request.Method == "OPTIONS" {
		// 	c.AbortWithStatus(200)
		// 	return
		// }

		// Continuar con el siguiente middleware
		c.Next()
	}
}
