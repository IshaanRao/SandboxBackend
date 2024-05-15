package main

import (
	"github.com/gin-gonic/gin"
	"sandboxbackend/middleware"
	"sandboxbackend/routes"
)

func main() {
	router := setupRouter()
	router.Run(":5600") // Default port is 8080
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	// Apply the API key middleware globally
	router.Use(middleware.APIKeyAuth())

	// Register routes
	router.GET("/players/:uuid", routes.GetPlayer)
	router.POST("/players/setrank/:uuid", routes.SetRank)
	router.GET("/servers/list", routes.ListServers)
	router.POST("/servers/proxyready", routes.ProxyReady)
	router.GET("/players/inv/:uuid", routes.GetInvContents)
	router.POST("/players/setinvcontents/:uuid", routes.SetInvContents)
	router.POST("/players/setarmorcontents/:uuid", routes.SetArmorContents)

	return router
}
