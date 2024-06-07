package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/words-reminder-api/api/models"
	"github.com/words-reminder-api/db/queries"
	"github.com/words-reminder-api/internal/provider"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var code int = http.StatusBadRequest
	var badResponse models.BadResponse
	var container provider.ContainerRep = provider.GETContainer()

	db, err := container.MYSQLConnector().Open(c)
	if err != nil {
		log.Warn(err, " | error openning db connection")
		badResponse.Code = code
		badResponse.Error = "error openning db connection"
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}
	defer db.Close()

	log.Info("Success db connection!")

	var u models.User

	if err := c.BindJSON(&u); err != nil {
		log.Warn(err, " | invalid json")
		badResponse.Code = code
		badResponse.Error = "invalid json"
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}

	if u.Email == "" || u.Password == "" {
		badResponse.Code = code
		badResponse.Error = "email and password required"
		log.Warn(badResponse.Error)
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}

	// buscar usuario en db
	q := queries.SQLFindUser

	var created []uint8

	var uf models.User
	// fmt.Println(uf)

	fmt.Println("query", q)

	err = db.QueryRowContext(c, q, u.Email).Scan(
		&uf.ID,
		&uf.Email,
		&uf.Password,
		&created,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Warn(err, " | user not found")
			badResponse.Code = http.StatusNotFound
			badResponse.Error = "user not found"
			c.IndentedJSON(http.StatusNotFound, badResponse)
			return
		}

		log.Warn(err, " | executing query")
		badResponse.Code = code
		badResponse.Error = "executing query"
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}

	// var layout string = "2023-04-08T20:48:01"

	fmt.Println(string(created))

	createdTime, err := time.Parse(time.DateTime, string(created))
	if err != nil {
		log.Warn(err, " | an error parsing date of creation")
		badResponse.Code = code
		badResponse.Error = "an error parsing date of creation"
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}

	uf.Created = createdTime

	// fmt.Print(uf)

	// comparar hash db con password recibida en request

	// respuesta en caso de exito

	// fmt.Println(uf.Password)
	// fmt.Println(u.Password)

	isPaswwordCheck := checkPasswordHash(u.Password, uf.Password)

	if !isPaswwordCheck {
		log.Warn(err, " | incorrect password")
		badResponse.Code = code
		badResponse.Error = "incorrect password"
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}

	code = http.StatusOK

	atk, err := models.CreateToken(uf.ID, uf.Email)

	if err != nil {
		log.Warn(err, " | creating access token")
		badResponse.Code = code
		badResponse.Error = "creating access token"
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}

	var ur models.LoginResponse
	ur.ID = uf.ID
	ur.Email = uf.Email
	ur.Created = uf.Created
	ur.AccessToken = atk

	var response models.Response
	response.Code = code
	response.Data = ur

	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	// c.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
	// c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	c.IndentedJSON(http.StatusOK, response)
}

func Logout(c *gin.Context) {

}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
