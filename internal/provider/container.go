package provider

import (
	"github.com/words-reminder-api/api/models"
	"github.com/words-reminder-api/db"
	"github.com/words-reminder-api/db/mysql"
)

var (
	container      ContainerRep   = nil
	mysqlconnector db.DataOpener  = nil
	user           models.UserRep = nil
	word           models.WordRep = nil
)

type ContainerRep interface {
	MYSQLConnector() db.DataOpener
	User() models.UserRep
	Word() models.WordRep
}

type Container struct{}

func GETContainer() ContainerRep {
	if container == nil {
		container = &Container{}
	}
	return container
}

func (c Container) MYSQLConnector() db.DataOpener {
	if mysqlconnector == nil {
		mysqlconnector = mysql.Connector{}
	}
	return mysqlconnector
}

func (c Container) User() models.UserRep {
	if user == nil {
		user = &models.User{}
	}
	return user
}

func (c Container) Word() models.WordRep {
	if word == nil {
		word = &models.Word{}
	}
	return word
}
