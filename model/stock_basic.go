package model

import (
	"github.com/guregu/null"
	"time"
)

// StockBasic table comment
type StockBasic struct {
	// BlacklistFlag 是否黑名单
	BlacklistFlag string `gorm:"column:blacklist_flag" json:"blacklist_flag"`
	// CalcuteDate 计算周涨跌幅完成时间,如果无效时间,表示不可交易
	CalcuteDate null.Time `gorm:"column:calcute_date" json:"calcute_date"`
	// CompanyTags 护城河,优质股,潜力股,十倍股,待观察
	CompanyTags string `gorm:"column:company_tags" json:"company_tags"`
	// CompareYesLastWeek 昨日收盘对比上周最高点涨跌幅0.1表示涨了10%,-0.1表示跌了10%
	CompareYesLastWeek float64 `gorm:"column:compare_yes_last_week" json:"compare_yes_last_week"`
	// ConID IB股票标记值
	ConID int `gorm:"column:con_id" json:"con_id"`
	// CreateTime column comments
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	// Currency USD
	Currency string `gorm:"column:currency" json:"currency"`
	// CurrentPrice 最新价
	CurrentPrice float64 `gorm:"column:current_price" json:"current_price"`
	// Description 描述
	Description string `gorm:"column:description" json:"description"`
	// Enabled 是否可用
	Enabled string `gorm:"column:enabled" json:"enabled"`
	// Exchange ISLAND/SMART
	Exchange string `gorm:"column:exchange" json:"exchange"`
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
	// Industry 所属行业
	Industry string `gorm:"column:industry" json:"industry"`
	// IsChina 中概股
	IsChina string `gorm:"column:is_china" json:"is_china"`
	// IsEtf 期权
	IsEtf string `gorm:"column:is_etf" json:"is_etf"`
	// IsHot 热门
	IsHot string `gorm:"column:is_hot" json:"is_hot"`
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
	// PrimExchange NASDAQ.NMS/NYSE
	PrimExchange string `gorm:"column:prim_exchange" json:"prim_exchange"`
	// SecType STK/OPT
	SecType string `gorm:"column:sec_type" json:"sec_type"`
	// Symbol 代码
	Symbol string `gorm:"column:symbol" json:"symbol"`
	// TotalMarketCap 总市值
	TotalMarketCap float64 `gorm:"column:total_market_cap" json:"total_market_cap"`
	// TradeDate 股票交易日
	TradeDate null.Time `gorm:"column:trade_date" json:"trade_date"`
	// TradeFlag 平稳向上,波动向上,高风险
	TradeFlag string `gorm:"column:trade_flag" json:"trade_flag"`
	// YesterdayPrice 昨收
	YesterdayPrice float64 `gorm:"column:yesterday_price" json:"yesterday_price"`
}

// TableName sets the insert table name for this struct type
func (s *StockBasic) TableName() string {
	return "stock_basic"
}
