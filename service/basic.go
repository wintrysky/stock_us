package service

import (
	"github.com/devfeel/mapper"
	"github.com/gin-gonic/gin"
	"github.com/wintrysky/orm"
	"github.com/wintrysky/stock_us/model"
	"github.com/wintrysky/stock_us/utils"
	"github.com/wintrysky/stock_us/utils/csv"
	"github.com/wintrysky/stock_us/xerr"
	"time"
)

func ImportBasicData(ctx *gin.Context,importType string) (err error){
	defer xerr.HandleErr(&err)

	filePath := utils.LoadExcelToTempFolder(ctx)

	cls := &csv.FileCsvService{}
	data,err := cls.ImportCSV(filePath)
	xerr.ThrowError(err)

	var tempList []model.StockBasic
	dt := time.Now()
	for _,value := range data {
		var tmp model.StockBasic
		mapper.MapperMap(value,&tmp)
		tmp.CreateTime = dt
		tempList = append(tempList,tmp)
	}

	err = refreshBasicData(tempList)
	return err
}

func refreshBasicData(tempList []model.StockBasic) (err error){
	var updateCols []string
	updateCols = append(updateCols,"current_price")
	updateCols = append(updateCols,"high_price")
	updateCols = append(updateCols,"increase_rate_10day")
	updateCols = append(updateCols,"increase_rate_120day")
	updateCols = append(updateCols,"increase_rate_20day")
	updateCols = append(updateCols,"increase_rate_250day")
	updateCols = append(updateCols,"increase_rate_5day")
	updateCols = append(updateCols,"increase_rate_60day")
	updateCols = append(updateCols,"increase_rate_form_year")
	updateCols = append(updateCols,"increase_rate_yesterday")
	updateCols = append(updateCols,"low_price")
	updateCols = append(updateCols,"open_price")
	updateCols = append(updateCols,"pe")
	updateCols = append(updateCols,"name")
	updateCols = append(updateCols,"total_market_cap")
	updateCols = append(updateCols,"trade_date")
	updateCols = append(updateCols,"yesterday_price")

	qa := orm.NewSession()
	qa.BatchUpsert(&tempList,500,updateCols,[]string{"symbol"})

	return qa.Error
}
