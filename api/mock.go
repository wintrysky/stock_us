package api

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/wintrysky/orm"
	"github.com/wintrysky/stock_us/model"
	"github.com/wintrysky/stock_us/utils"
)

func Demo(ctx *gin.Context) {
	var item model.UnitTestModel

	qa := orm.NewSession()
	qa.AutoMigrate(&item)
	log.Info("-------")
	utils.ApiErrorResponse(ctx,"ERROR TESt")

}