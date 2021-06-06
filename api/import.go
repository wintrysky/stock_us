package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wintrysky/stock_us/service"
	"github.com/wintrysky/stock_us/utils"
	"net/url"
)

// ImportCSV 导入CSV文件
func ImportCSV(ctx *gin.Context) {
	// nginx下不能用下划线
	importType := ctx.Request.Header.Get("importtype")
	if importType == "" {
		utils.ApiErrorResponse(ctx,"没有导入类型")
		return
	}
	importTypeX, _ := url.QueryUnescape(importType)

	var err error
	if importTypeX == "all" {
		err = service.ImportBasicData(ctx,importTypeX)
	}

	utils.ApiResponse(ctx,nil,err)
}