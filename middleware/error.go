package middleware

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/wintrysky/stock_us/global"
	"io"
	"net/http"
)

// ErrorHandle 统一捕获错误
func ErrorHandle(out io.Writer) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				abortWithError(ctx,err)
			}
		}()
		ctx.Next()
	}
}

func abortWithError(ctx *gin.Context, err interface{}){
	log.Error(err)
	ctx.JSON(http.StatusOK, &global.Response{Code: -1, Data: cast.ToString(err)})
}
