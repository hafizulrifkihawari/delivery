package controllers

import (
	"delivery/response"

	"github.com/gin-gonic/gin"
)

// Health check used to confirm service status
func HealthCheck(c *gin.Context) {
	var resp response.IResponse = response.Response{}
	resp.Success(c)
}
