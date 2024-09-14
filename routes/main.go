package routes

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	var router *gin.Engine = gin.Default()
	getRoutes(router)
	return router
}

func getRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	addPingRoutes(v1)
}

func Run() {
	var router *gin.Engine = setupRouter()
	router.Run(":5000")
}
