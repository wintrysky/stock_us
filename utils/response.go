package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/wintrysky/stock_us/global"
	"net/http"
)

func ApiResponse(ctx *gin.Context,rst interface{},err error){
	if err != nil {
		ctx.JSON(http.StatusOK, &global.Response{Code: -1, Message: err.Error()})
	}else{
		ctx.JSON(http.StatusOK, &global.Response{Code: 0, Data: rst})
	}
}

func ApiErrorResponse(ctx *gin.Context,errMsg string){
	ctx.JSON(http.StatusOK, &global.Response{Code: -1, Message: errMsg})
}