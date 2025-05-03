package controllers

import (
	"net/http"

	"github.com/emmanueljustin/go-api/services"
	"github.com/gin-gonic/gin"
)

func GetSampleData(c *gin.Context) {
	data, err := services.GetSampleData()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": data,
		"status":  "success",
	})
}
