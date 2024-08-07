package routes

import (
	"DumbiFadhil/edas-api/controllers"
	"DumbiFadhil/edas-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	apiV1 := router.Group("/api/v1")
	{
		apiV1.POST("/edas", controllers.CalculateEDAS)

		apiV1.GET("/history", func(c *gin.Context) {
			history, err := services.GetAllHistory()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to get history"})
				return
			}
			c.JSON(http.StatusOK, history)
		})

		apiV1.DELETE("/history/:uuid", func(c *gin.Context) {
			uuid := c.Param("uuid")
			err := services.DeleteHistoryByUUID(uuid)
			if err != nil {
				if err != nil {
					c.JSON(http.StatusNotFound, gin.H{"status": "error", "message": "History not found"})
					return
				}
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to delete history"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "success", "message": "History successfully deleted"})
		})

		apiV1.GET("/history/:uuid", func(c *gin.Context) {
			uuid := c.Param("uuid")
			history, err := services.GetHistoryByUUID(uuid)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Failed to get history"})
				return
			}
			c.JSON(http.StatusOK, history)
		})
	}
}
