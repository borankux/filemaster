package main

import (
	"github.com/borankux/filemaster/server/config"
	"github.com/borankux/filemaster/server/db"
	"github.com/borankux/filemaster/server/router"
)

func main() {
	config.Init("dev")
	db.Init()
	router.Init()
}
