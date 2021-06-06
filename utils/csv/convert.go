package csv

import (
	"github.com/spf13/cast"
	"github.com/wintrysky/stock_us/utils"
	"github.com/wintrysky/stock_us/xerr"
	"strings"
	"time"
)

func toString(input string) string {
	return utils.TrimSpace(input)
}

func toFloat(input string) float64 {
	tmp := cast.ToFloat64(input)
	if tmp == 0 {
		tmp = -9999
	}
	return tmp
}

func toMoney(input string) float64 {
	var tmpF float64
	if strings.Contains(input,"亿"){
		tmp := strings.ReplaceAll(input,"亿","")
		tmpF = cast.ToFloat64(tmp)
	}else if strings.Contains(input,"万"){
		tmp := strings.ReplaceAll(input,"万","")
		tmpV := cast.ToFloat64(tmp)
		tmpF =  tmpV / 10000 // 万
	}else{
		tmpV := cast.ToFloat64(input) // 如果没有到达万，就是个位数
		tmpF =  tmpV / (10000*10000) // 万
	}
	return tmpF
}

func toPercent(input string) float64 {
	removeFlag := strings.ReplaceAll(input,"%","")
	tmp := cast.ToFloat64(removeFlag)
	if tmp == 0 {
		tmp = -9999
	}
	return tmp
}

func toInt(input string) int64 {
	tmp := cast.ToInt64(input)
	if tmp == 0 {
		tmp = -9999
	}
	return tmp
}

func toDatetime(input string) time.Time {
	tmp,err := cast.ToTimeE(input)
	xerr.ThrowError(err)

	return tmp
}