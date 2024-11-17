package drivers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/api/models"
	"github.com/words-reminder-api/helpers"
	"github.com/words-reminder-api/internal/provider"
)

type GPTResp struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Choices []Choice `json:"choices"`
}

type Choice struct {
	Index   int `json:"index"`
	Message `json:"message"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func CreateNewWord(c *gin.Context, newWord *models.Word) error {

	var container provider.ContainerRep = provider.GETContainer()
	var f models.WordFilter
	f.Content = newWord.Content

	fmt.Println("searching word...", f.Content)
	word, err := container.Word().Search(c, f)
	if err != nil {
		if err != sql.ErrNoRows {
			err = errors.New("executing Search: " + err.Error())
			return err
		}
	}

	if word != nil && word.ID > 0 {
		f.UserID = newWord.UserID
		f.WordID = word.ID
		wordSearched, err := container.Word().SearchByUserID(c, f)
		if err != nil {
			err = errors.New("executing SearchByUserID: " + err.Error())
			return err
		}

		if wordSearched.ID == 0 {
			err = container.Word().Link(c, newWord.UserID, word.ID)
			if err != nil {
				err = errors.New("executing Link: " + err.Error())
				return err
			}
			return nil
		} else {
			return errors.New("word already exists for this user")
		}
	}

	apiURL := "https://api.openai.com/v1/chat/completions"
	apiKey := os.Getenv("GPT_API_KEY")

	requestBody := map[string]interface{}{
		"model": "gpt-4o-mini",
		"messages": []map[string]string{
			{"role": "user", "content": "Traduce la siguiente palabra al español: " + f.Content + ". Importante: Retorna solo la palabra traducida."},
		},
		"max_tokens":  50,
		"temperature": 0.1,
	}

	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer " + apiKey,
	}

	res, err := helpers.MakeHTTPRequest("POST", apiURL, headers, requestBody)
	if err != nil {
		return errors.New("error executing MakeHTTPRequest: " + err.Error())
	}

	var gr GPTResp
	err = json.Unmarshal(res, &gr)
	if err != nil {
		return err
	}

	newWord.Content = strings.ToLower(newWord.Content)
	translation := strings.ReplaceAll(gr.Choices[0].Content, ".", "")
	newWord.Translation = strings.ToLower(translation)

	err = container.Word().Create(c, newWord)
	if err != nil {
		return err
	}

	word, err = container.Word().Search(c, f)
	if err != nil {
		return err
	}

	err = container.Word().Link(c, newWord.UserID, word.ID)
	if err != nil {
		return err
	}

	return nil
}

func FetchWords(c *gin.Context, f models.WordFilter) ([]*models.Word, error) {
	var container provider.ContainerRep = provider.GETContainer()

	words, err := container.Word().Fetch(c, f)
	if err != nil {
		return words, err
	}
	return words, nil
}
