package model

import (
	"time"
)

// StockBasicHistory table comment
type StockBasicHistory struct {
	// CompareYesLastWeek 昨日收盘对比上周最高点涨跌幅0.1表示涨了10%,-0.1表示跌了10%
	CompareYesLastWeek float64 `gorm:"column:compare_yes_last_week" json:"compare_yes_last_week"`
	// CurrentPrice 最新价
	CurrentPrice float64 `gorm:"column:current_price" json:"current_price"`
	// HighPrice 最高
	HighPrice float64 `gorm:"column:high_price" json:"high_price"`
	// ID column comments
	ID int64 `gorm:"primary_key;column:id" json:"id"`
	// IncreaseRate10day 10日涨跌幅
	IncreaseRate10day float64 `gorm:"column:increase_rate_10day" json:"increase_rate_10day"`
	// IncreaseRate120day 120日涨跌幅
	IncreaseRate120day float64 `gorm:"column:increase_rate_120day" json:"increase_rate_120day"`
	// IncreaseRate20day 20日涨跌幅
	IncreaseRate20day float64 `gorm:"column:increase_rate_20day" json:"increase_rate_20day"`
	// IncreaseRate250day 250日涨跌幅
	IncreaseRate250day float64 `gorm:"column:increase_rate_250day" json:"increase_rate_250day"`
	// IncreaseRate5day 5日涨跌幅
	IncreaseRate5day float64 `gorm:"column:increase_rate_5day" json:"increase_rate_5day"`
	// IncreaseRate60day 60日涨跌幅
	IncreaseRate60day float64 `gorm:"column:increase_rate_60day" json:"increase_rate_60day"`
	// IncreaseRateFormYear 年初至今涨跌幅
	IncreaseRateFormYear float64 `gorm:"column:increase_rate_form_year" json:"increase_rate_form_year"`
	// IncreaseRateYesterday 涨跌幅
	IncreaseRateYesterday float64 `gorm:"column:increase_rate_yesterday" json:"increase_rate_yesterday"`
	// LastWeekRate 上周涨跌率,0.1表示涨了10%,-0.1表示跌了10%
	LastWeekRate float64 `gorm:"column:last_week_rate" json:"last_week_rate"`
	// LowPrice 最低
	LowPrice float64 `gorm:"column:low_price" json:"low_price"`
	// Name 名称
	Name string `gorm:"column:name" json:"name"`
	// OpenPrice 开盘
	OpenPrice float64 `gorm:"column:open_price" json:"open_price"`
	// Pe 市盈率(静)
	Pe float64 `gorm:"column:pe" json:"pe"`
	// Symbol 代码
	Symbol string `gorm:"column:symbol" json:"symbol"`
	// TotalMarketCap 总市值
	TotalMarketCap float64 `gorm:"column:total_market_cap" json:"total_market_cap"`
	// TradeDate 交易日
	TradeDate time.Time `gorm:"column:trade_date" json:"trade_date"`
	// YesterdayPrice 昨收
	YesterdayPrice float64 `gorm:"column:yesterday_price" json:"yesterday_price"`
}

// TableName sets the insert table name for this struct type
func (s *StockBasicHistory) TableName() string {
	return "stock_basic_history"
}
