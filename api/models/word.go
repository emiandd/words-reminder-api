package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/db/queries"
)

type Word struct {
	ID            int64     `json:"id"`
	Content       string    `json:"content"`
	Status        int       `json:"status"`
	UserID        int64     `json:"user_id"`
	YouGlish      string    `json:"youglish"`
	Created       time.Time `json:"created"`
	Translation   string    `json:"translation"`
	AlreadyExists bool      `json:"already_exists"`
}

type WordFilter struct {
	UserID  int64  `form:"user_id"`
	WordID  int64  `form:"word_id"`
	Content string `form:"content"`
}

type WordRep interface {
	Create(c *gin.Context, data *Word) error
	Fetch(c *gin.Context, f WordFilter) ([]*Word, error)
	Search(c *gin.Context, f WordFilter) (*Word, error)
	SearchByUserID(c *gin.Context, f WordFilter) (*Word, error)
	Link(c *gin.Context, userID, wordID int64) error
	Update(c *gin.Context) error
}

var (
	_ WordRep = &Word{}
)

func (w *Word) Create(c *gin.Context, data *Word) error {
	db, ok := (c.Value("conn")).(*sql.DB)
	if !ok {
		return fmt.Errorf("error obtaining connection to database: %v", db)
	}

	q := queries.SQLCreateNewWord

	result, err := db.Exec(q, data.Content, data.Status, data.UserID, data.YouGlish, time.Now(), data.Translation)
	if err != nil {
		err = fmt.Errorf("error executing query: %s \n %v", q, err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		err = fmt.Errorf("error getting last inserted id: %v", err)
		return err
	}

	w.ID = id

	return nil
}

func (w *Word) Fetch(c *gin.Context, f WordFilter) ([]*Word, error) {
	db, ok := (c.Value("conn")).(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("error obtaining connection to database: %v", db)
	}

	q, params := fetchQueryBuilder(f)

	// q := queries.SQLFetchWords

	fmt.Println("query: ", q)
	fmt.Println("params: ", params)

	rows, err := db.Query(q, params...)
	if err != nil {
		err = fmt.Errorf("error executing query: %s \n %v", q, err)
		return nil, err
	}

	var words []*Word
	for rows.Next() {
		var word Word
		var created []uint8
		err = rows.Scan(
			&word.ID,
			&word.Content,
			&word.Status,
			&word.YouGlish,
			&created,
			&word.Translation,
		)

		if err != nil {
			return nil, err
		}

		createdTime, err := time.Parse(time.DateTime, string(created))
		if err != nil {
			return nil, err
		}

		word.Created = createdTime

		words = append(words, &word)
	}

	return words, nil
}

func (w *Word) Search(c *gin.Context, f WordFilter) (*Word, error) {
	db, ok := (c.Value("conn")).(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("error obtaining connection to database: %v", db)
	}

	q := queries.SQLSearchWord

	var word Word
	err := db.QueryRow(q, f.Content).Scan(
		&word.ID,
		&word.Content,
	)

	if err != nil {
		err = fmt.Errorf("error executing query: %s \n %v", q, err)
		return nil, err
	}

	return &word, nil
}

func (w *Word) SearchByUserID(c *gin.Context, f WordFilter) (*Word, error) {
	db, ok := (c.Value("conn")).(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("error obtaining connection to database: %v", db)
	}

	q := queries.SQLSearchWordByUser

	fmt.Println("query: ", q)
	fmt.Println("params", f.UserID, f.WordID)

	var word Word
	err := db.QueryRow(q, f.UserID, f.WordID).Scan(
		&word.ID,
		&word.Content,
	)

	if err != nil {
		if err != sql.ErrNoRows {
			err = fmt.Errorf("error executing query: %s \n %v", q, err)
			return nil, err
		}

	}

	return &word, nil
}

func (w *Word) Link(c *gin.Context, userID, wordID int64) error {
	db, ok := (c.Value("conn")).(*sql.DB)
	if !ok {
		return fmt.Errorf("error obtaining connection to database: %v", db)
	}

	q := queries.SQLLinkWordToUser

	result, err := db.Exec(q, userID, wordID)
	if err != nil {
		err = fmt.Errorf("error executing query: %s \n %v", q, err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		err = fmt.Errorf("error getting last inserted id: %v", err)
		return err
	}

	if id == 0 {
		return errors.New("something went wrong when trying to insert word_domain")
	}

	return nil
}

func (w *Word) Update(c *gin.Context) error {
	return nil
}
