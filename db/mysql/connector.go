package mysql

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/db"
)

var (
	_ db.DataOpener = Connector{}
)

type Connector struct{}

func (co Connector) Open(c *gin.Context) (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:wr_password@tcp(127.0.0.1:3306)/wr_development")
	if err != nil {
		return nil, err
	}

	// c.Keys = map[string]
	return db, nil
}
