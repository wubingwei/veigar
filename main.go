package main

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wubingwei/veigar/config"
	"github.com/wubingwei/veigar/handler"
)

func init() {
	if config.DevMode {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	router := gin.Default()
	handler.InitRouter(router)
	err := router.Run(":1994")
	if err != nil {
		log.Fatal(err)
	}
}
