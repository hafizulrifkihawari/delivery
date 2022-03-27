package gateway

import (
	"delivery/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
	defer func(c *gin.Context) {
		// panic recovery
		if rec := recover(); rec != nil {
			response := response.Response{Status: http.StatusInternalServerError}
			response.Error(c, http.StatusText(http.StatusInternalServerError))
		}
	}(c)
	c.Next()
}
