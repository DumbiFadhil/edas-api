package controllers

import (
	"DumbiFadhil/edas-api/models"
	"DumbiFadhil/edas-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CalculateEDAS(c *gin.Context) {
	var request models.EDASRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response := services.CalculateEDAS(request)
	c.JSON(http.StatusOK, response)
}
