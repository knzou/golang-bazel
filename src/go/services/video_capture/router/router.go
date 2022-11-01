package router

import (
	"monorepo/src/go/services/lib/handlers/health"

	"github.com/gin-gonic/gin"
)

func GetEngine() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.GET("/", health.Health("Video Capture app"))
	return r
}
