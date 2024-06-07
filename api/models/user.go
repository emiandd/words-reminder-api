package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/db/queries"
)

type User struct {
	ID       int64     `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password,omitempty"`
	Created  time.Time `json:"created"`
}

type UserFilter struct {
	Email string `form:"email"`
	FilterFetch
}

type UserRep interface {
	Create(c *gin.Context, hashedPassword string) error
	Fetch(c *gin.Context, f UserFilter) ([]*User, error)
	Count(c *gin.Context, f UserFilter) (int, error)
}

var (
	_ UserRep = &User{}
)

func (u *User) Create(c *gin.Context, hashedPassword string) error {
	db, ok := (c.Value("conn")).(*sql.DB)
	if !ok {
		return fmt.Errorf("error obtaining connection to database: %v", db)
	}

	q := queries.SQLCreateNewUser

	u.Created = time.Now()

	result, err := db.Exec(q, u.Email, hashedPassword, u.Created)
	if err != nil {
		err = fmt.Errorf("error executing query: %s \n %v", q, err)
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		err = fmt.Errorf("error getting last inserted id: %v", err)
		return err
	}

	u.ID = id

	return nil
}

func (u *User) Fetch(c *gin.Context, f UserFilter) ([]*User, error) {

	db, ok := (c.Value("conn")).(*sql.DB)
	if !ok {
		return nil, fmt.Errorf("error obtaining connection to database: %v", db)
	}

	q, params := userQueryBuilder(f)

	fmt.Println("query: ", q)
	fmt.Println("params: ", params)

	rows, err := db.Query(q, params...)
	if err != nil {
		return nil, err
	}

	var users []*User
	for rows.Next() {
		var user User
		var created []uint8
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Password,
			&created,
		)
		if err != nil {
			return nil, err
		}

		createdTime, err := time.Parse(time.DateTime, string(created))
		if err != nil {
			return nil, err
		}

		user.Created = createdTime

		users = append(users, &user)
	}

	return users, nil
}

func (u *User) Count(c *gin.Context, f UserFilter) (int, error) {
	db, ok := (c.Value("conn")).(*sql.DB)
	if !ok {
		return 0, fmt.Errorf("error obtaining connection to database: %v", db)
	}

	f.Count = true

	q, params := userQueryBuilder(f)

	fmt.Println("query: ", q)
	fmt.Println("params: ", params)

	var count int
	err := db.QueryRow(q, params...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("error scan user count: %v", err)
	}

	return count, nil
}
