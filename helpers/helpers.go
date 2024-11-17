package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EndpointPagination(c *gin.Context, count, limit, offset int) (prev string, next string) {
	url := "http://localhost:8080" + c.FullPath()
	// @TODO i0 - soportar query params - Fecha: 18 April, 2024

	if offset == 0 {
		next = fmt.Sprintf("%s?offset=%d", url, limit+offset)
	}

	if offset > 0 {
		if limit-offset == 0 {
			prev = url
		} else {
			prev = fmt.Sprintf("%s?offset=%v", url, math.Abs(float64(limit-offset)))
		}

		if offset+limit > count {
			next = ""
		} else {
			next = fmt.Sprintf("%s?offset=%d", url, limit+offset)
		}
	}

	return prev, next
}

func MakeHTTPRequest(method, url string, headers map[string]string, body interface{}) ([]byte, error) {
	var requestBody []byte
	var err error
	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("error al serializar el cuerpo de la solicitud: %v", err)
		}
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("error al crear la solicitud HTTP: %v", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error al realizar la solicitud HTTP: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error al leer el cuerpo de la respuesta: %v", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("error en la respuesta HTTP, código de estado: %d, cuerpo: %s", resp.StatusCode, string(responseBody))
	}

	return responseBody, nil
}
