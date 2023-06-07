package base

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type ginRouterSingleton struct {
	router *gin.Engine
	once   sync.Once
}

var ginRouterSingle = &ginRouterSingleton{}

func GinRouter() *gin.Engine {
	ginRouterSingle.once.Do(func() {
		ginRouterSingle.router = gin.Default()
	})
	return ginRouterSingle.router
}

func GinSimpleResponse(c *gin.Context, code uint32, msg string) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg})
}

func GinDataResponse(c *gin.Context, code uint32, msg string, data any) {
	c.JSON(http.StatusOK, gin.H{"code": code, "msg": msg, "data": data})
}
