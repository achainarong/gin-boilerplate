package server

import (
	"gin/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router:= gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// does'nt need to be versioned
	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	v1 := router.Group("api/v1")
	{
		hello := new(controllers.HelloController)
		v1.GET("/hello", hello.Hello)
	}

	router.NoRoute(func (context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})
	
	return router;
}