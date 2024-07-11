package config

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.New()

	// Attach middleware
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Configure trusted proxies if needed
	// router.SetTrustedProxies([]string{"192.168.1.1"})

	return router
}
