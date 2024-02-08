package controller

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	router := gin.Default()
	// router.LoadHTMLGlob("templates/*")
	return router
}
