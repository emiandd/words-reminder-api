package db

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type DataOpener interface {
	Open(c *gin.Context) (*sql.DB, error)
}
