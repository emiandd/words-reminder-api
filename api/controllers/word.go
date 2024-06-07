package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/words-reminder-api/api/drivers"
	"github.com/words-reminder-api/api/models"
)

func CreateNewWord(c *gin.Context) {
	var code int = http.StatusBadRequest
	var badResponse models.BadResponse
	var newWord models.Word

	if err := c.BindJSON(&newWord); err != nil {
		log.Warn(err, " | error binding word")
		badResponse.Code = code
		badResponse.Error = "error binding word"
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}

	fmt.Println("NEW WORD", newWord)

	if newWord.Content == "" || newWord.YouGlish == "" {
		badResponse.Code = code
		badResponse.Error = "content and youglish required"
		log.Warn(badResponse.Error)
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}

	err := drivers.CreateNewWord(c, &newWord)
	if err != nil {
		badResponse.Code = code
		badResponse.Error = "creating new word: " + err.Error()
		log.Warn(badResponse.Error)
		c.IndentedJSON(http.StatusBadRequest, badResponse)
		return
	}

	code = http.StatusCreated

	// newWord = ""

	var response models.Response
	response.Code = code
	response.Data = newWord

	c.IndentedJSON(http.StatusCreated, response)

}

func FetchWords(c *gin.Context) {
	var code int = http.StatusBadRequest
	var badResponse models.BadResponse
	var f models.WordFilter

	err := c.ShouldBindQuery(&f)
	if err != nil {
		log.Warn("obtain query params")
		badResponse.Code = http.StatusBadRequest
		badResponse.Error = "obtain query params"
		c.IndentedJSON(http.StatusUnauthorized, badResponse)
		return
	}

	fmt.Println("FILTER FETCH: ", f)

	words, err := drivers.FetchWords(c, f)
	if err != nil {
		log.Warn("fetching words" + err.Error())
		badResponse.Code = http.StatusBadRequest
		badResponse.Error = "fetching words"
		c.IndentedJSON(http.StatusUnauthorized, badResponse)
		return
	}

	code = http.StatusOK

	var rm models.ResponseMany
	rm.Code = code
	rm.Data = words
	rm.Count = 0
	rm.Prev = "prev"
	rm.Next = "next"

	c.IndentedJSON(http.StatusCreated, rm)

}
