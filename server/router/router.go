package router

import (
	"github.com/borankux/filemaster/server/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		location := url.URL{Path: "/health"}
		c.Redirect(http.StatusFound, location.RequestURI())
	})
	router.GET("health", controllers.HealthController{}.Status)
	return router
}
