package main

import (
	"github.com/joyllee/blocks/config"
	"github.com/joyllee/blocks/logger"
)

func main() {
	config.LoadConfig("./config/dev.yaml")
	logger.InitLogger()

	logger.Infof("init !!!!")
}
