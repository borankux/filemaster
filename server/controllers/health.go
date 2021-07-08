package controllers

import (
	"github.com/borankux/filemaster/server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HealthController struct {

}

func(h HealthController) Status(c *gin.Context) {
	r :=models.Request{
		Ip: c.ClientIP(),
	}.Save()

	c.JSON(http.StatusOK, &r)
}