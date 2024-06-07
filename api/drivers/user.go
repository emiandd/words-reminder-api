package drivers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/api/models"
	"github.com/words-reminder-api/internal/provider"
	"golang.org/x/crypto/bcrypt"
)

func CreateNewUser(c *gin.Context, newUser *models.User) error {

	hashedPassword, err := HashPassword(newUser.Password)
	if err != nil {
		err = fmt.Errorf("HashPassword: %v", err)
		return err
	}

	var container provider.ContainerRep = provider.GETContainer()

	err = container.User().Create(c, hashedPassword)

	// err = newUser.Create(c, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func FetchUsers(c *gin.Context, f models.UserFilter) (users []*models.User, err error) {
	var container provider.ContainerRep = provider.GETContainer()

	users, err = container.User().Fetch(c, f)
	if err != nil {
		return users, err
	}

	return users, nil
}

func CountUsers(c *gin.Context, f models.UserFilter) (int, error) {
	var container provider.ContainerRep = provider.GETContainer()

	count, err := container.User().Count(c, f)
	if err != nil {
		return count, err
	}

	return count, nil
}