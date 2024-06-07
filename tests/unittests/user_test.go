package unit

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/words-reminder-api/api/models"
)

// TestPOST: Create a new user
func TestPOST(t *testing.T) {
	url := "http://localhost:8080/v1/users"
	method := http.MethodPost

	payload := strings.NewReader(`
	{
    "email": "correox@correop.com",
    "password": "testpass"
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	res, err := client.Do(req)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	var ur models.User
	err = json.Unmarshal(body, &ur)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	fmt.Println(string(body))
}

// TestGetAllUsers: obtain all users
func TestGetAllUsers(t *testing.T) {
	url := "http://localhost:8080/v1/users"
	method := http.MethodGet

	payload := strings.NewReader(`
	{
    "access_token": ""
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	res, err := client.Do(req)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Log(err)
		t.FailNow()
	}

	var users []models.User
	err = json.Unmarshal(body, &users)

	if err != nil {
		t.Log(err)
		t.FailNow()
	}
	fmt.Println(string(body))
}
