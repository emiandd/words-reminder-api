package drivers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/words-reminder-api/api/models"
	"github.com/words-reminder-api/internal/provider"
)

func CreateNewWord(c *gin.Context, newWord *models.Word) error {

	var container provider.ContainerRep = provider.GETContainer()
	var f models.WordFilter
	f.Content = newWord.Content
	// f.UserID = newWord.UserID

	fmt.Println("searching word...", f.Content)
	word, err := container.Word().Search(c, f)
	if err != nil {
		err = errors.New("executing Search: " + err.Error())

		return err
	}

	fmt.Println("WORD: ", word)
	if word.ID > 0 {

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

	url := "https://microsoft-translator-text.p.rapidapi.com/translate?api-version=3.0&to%5B0%5D=es&textType=plain&profanityAction=NoAction"
	payload := strings.NewReader(fmt.Sprintf("[\r\n    {\r\n        \"Text\": \"%s\"\r\n    }\r\n]", newWord.Content))
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-RapidAPI-Key", "673c914583msh8abafbd0616306ap11b024jsn6853dd934fb1")
	req.Header.Add("X-RapidAPI-Host", "microsoft-translator-text.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	respTranslatorApi := make([]interface{}, 1)

	err = json.NewDecoder(res.Body).Decode(&respTranslatorApi)
	if err != nil {
		return err
	}

	dataMap, ok := (respTranslatorApi[0]).(map[string]interface{})
	if !ok {
		return errors.New("respTranslatorApi assertion type failed")
	}

	translations := dataMap["translations"].([]interface{})
	translationsMap := translations[0].(map[string]interface{})
	newWord.Translation = strings.ToLower(translationsMap["text"].(string))

	err = container.Word().Create(c, newWord)
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
