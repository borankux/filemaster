package router

import (
	"github.com/borankux/filemaster/server/config"
	"log"
)

func Init()  {
	r := NewRouter()
	conf := config.GetConf()
	err := r.Run(":"+conf.GetString("server.port"))
	if err != nil {
		log.Fatalf("application failed to start:%v", err)
	}
}