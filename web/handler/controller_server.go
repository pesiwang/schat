package handler

import (
	"schat/web/base"
	"schat/web/define"

	"github.com/gin-gonic/gin"
)

func initControllerServer() {
	serverGroup := base.GinRouter().Group("/server")

	serverGroup.GET("/health_check", healthCheck)
}

func healthCheck(c *gin.Context) {
	base.GinSimpleResponse(c, define.RESULT_CODE_SUCC, "succ")
}
