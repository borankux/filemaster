package main

import (
	"github.com/borankux/filemaster/server/config"
	"github.com/borankux/filemaster/server/db"
	"github.com/borankux/filemaster/server/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	config.Init("dev")
	db.Init()
	router.Init()
}
