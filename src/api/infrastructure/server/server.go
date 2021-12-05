package server

import "github.com/gin-gonic/gin"

func Start() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())
	mapRoutes(router)
	router.Run(":9081")
}