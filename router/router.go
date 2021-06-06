package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wintrysky/stock_us/api"
	"github.com/wintrysky/stock_us/middleware"
	"io"
	"os"
)

// UseRouters 定义路由
func UseRouters(eng *gin.Engine) {
	rg := eng.Group("/")
	var errWriter io.Writer = os.Stderr
	rg.Use(middleware.ErrorHandle(errWriter))
	useBizRouters(rg)
}

// 定义业务路由
func useBizRouters(rg *gin.RouterGroup) {
	rg.GET("/demo", api.Demo)
	rg.POST("/import", api.ImportCSV)
}
