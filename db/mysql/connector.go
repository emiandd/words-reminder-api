package mysql

import (
	"database/sql"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/db"
)

var (
	_ db.DataOpener = Connector{}
)

type Connector struct{}

func (co Connector) Open(c *gin.Context) (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("USER_DB")+":" + os.Getenv("PASS_DB") + "@tcp("+ os.Getenv("HOST_DB") +":"+ os.Getenv("PORT_DB") +")/"+os.Getenv("NAME_DB")+"?parseTime=true")
	if err != nil {
		return nil, err
	}

	return db, nil
}
