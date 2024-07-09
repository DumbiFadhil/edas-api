package routes

import (
    "DumbiFadhil/edas-api/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.POST("/api/edas", controllers.CalculateEDAS)
}

