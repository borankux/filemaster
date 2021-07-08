package main

import (
	"flag"
	"fmt"
	"github.com/borankux/filemaster/server/boot"
	"github.com/borankux/filemaster/server/config"
	"github.com/borankux/filemaster/server/db"
	"github.com/borankux/filemaster/server/router"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	env := flag.String("e", "dev", "")

	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parsed()

	gin.SetMode(gin.DebugMode)
	config.Init(*env)
	db.Init()
	boot.Migrate()
	router.Init()
}