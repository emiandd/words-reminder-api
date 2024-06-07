package helpers

import (
	"fmt"
	"math"

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
