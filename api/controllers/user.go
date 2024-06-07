package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	log "github.com/sirupsen/logrus"
	"github.com/words-reminder-api/api/drivers"
	"github.com/words-reminder-api/api/models"
	"github.com/words-reminder-api/helpers"
	// "golang.org/x/perf/storage/db"
)

func CreateNewUser(c *gin.Context) {
	var code int = http.StatusBadRequest
	var badResponse models.BadResponse
	var newUser models.User

	if err := c.BindJSON(&newUser); err != nil {
		log.Warn(err, " | error binding user")
		badResponse.Code = code
		badResponse.Error = "error binding user"
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}

	// fmt.Println("EMAIL", newUser.Email)
	// fmt.Println("PASSWORD", newUser.Password)

	if newUser.Email == "" || newUser.Password == "" {
		badResponse.Code = code
		badResponse.Error = "email and password required"
		log.Warn(badResponse.Error)
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}

	err := drivers.CreateNewUser(c, &newUser)
	if err != nil {
		badResponse.Code = code
		badResponse.Error = err.Error()
		log.Warn(badResponse.Error)
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}

	code = http.StatusCreated

	newUser.Password = ""

	var response models.Response
	response.Code = code
	response.Data = newUser

	c.IndentedJSON(http.StatusCreated, response)
}

func GetUsers(c *gin.Context) {
	var code int = http.StatusBadRequest
	var badResponse models.BadResponse
	var f models.UserFilter

	err := c.ShouldBindQuery(&f)
	if err != nil {
		log.Warn("obtain query params")
		badResponse.Code = http.StatusBadRequest
		badResponse.Error = "obtain query params"
		c.IndentedJSON(http.StatusUnauthorized, badResponse)
		return
	}

	if f.Limit == 0 {
		f.Limit = 25
	}

	fmt.Println("LIMIT:", f.Limit)

	// f.Limit = limit
	// f.Offset = offset
	count, err := drivers.CountUsers(c, f)
	if err != nil {
		log.Warn(err, " | obtaining users count")
		badResponse.Code = http.StatusBadRequest
		badResponse.Error = "obtaining users count"
		c.IndentedJSON(http.StatusUnauthorized, badResponse)
		return
	}

	users, err := drivers.FetchUsers(c, f)
	if err != nil {
		log.Warn(err, " | fetching users")
		badResponse.Code = http.StatusBadRequest
		badResponse.Error = "fetching users"
		c.IndentedJSON(http.StatusUnauthorized, badResponse)
		return
	}

	code = http.StatusOK
	prev, next := helpers.EndpointPagination(c, count, f.Limit, f.Offset)

	var rm models.ResponseMany
	rm.Code = code
	rm.Data = users
	rm.Count = count
	rm.Prev = prev
	rm.Next = next

	c.IndentedJSON(http.StatusOK, rm)
}
