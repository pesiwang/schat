package main

import (
	"schat/infrastructure/logger"
	"schat/web/base"
	"schat/web/handler"
)

func main() {
	logger.InitLogger("/dev/stdout")
	handler.InitHandler()
	base.GinRouter().Run(":8888")
}
