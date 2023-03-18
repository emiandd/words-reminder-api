package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, albums)
}

func CreateNewUser(c *gin.Context) {

	db, err := database.DBConnector()

	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	fmt.Print(db)

	var newUser User

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	fmt.Print(newUser)

	// Add the new album to the slice.
	// albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newUser)
}
