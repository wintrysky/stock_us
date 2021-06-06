/*
SQLyog 企业版 - MySQL GUI v8.12,3 
MySQL - 5.7.18-20170830-log 
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `stock_basic_history` (
	`id` bigint NOT NULL AUTO_INCREMENT,
	`symbol` varchar (100) comment '代码',
	`name` varchar (200) comment '名称',
	`open_price` decimal (12,3) comment '开盘',
	`low_price` decimal (12,3) comment '最低',
	`high_price` decimal (12,3) comment '最高',
	`current_price` decimal (12,3) comment '最新价',
	`yesterday_price` decimal (12,3) comment '昨收',
	`pe` decimal (12,3) comment '市盈率(静)',
	`total_market_cap` decimal (12,3) comment '总市值',
	`increase_rate_yesterday` decimal (12,3) comment '涨跌幅',
	`increase_rate_5day` decimal (12,3) comment '5日涨跌幅',
	`increase_rate_10day` decimal (12,3) comment '10日涨跌幅',
	`increase_rate_20day` decimal (12,3) comment '20日涨跌幅',
	`increase_rate_60day` decimal (12,3) comment '60日涨跌幅',
	`increase_rate_120day` decimal (12,3) comment '120日涨跌幅',
	`increase_rate_250day` decimal (12,3) comment '250日涨跌幅',
	`increase_rate_form_year` decimal (12,3) comment '年初至今涨跌幅',
	`trade_date` datetime comment '交易日',
	`last_week_rate` decimal (12,3) comment '上周涨跌率,0.1表示涨了10%,-0.1表示跌了10%',
	`compare_yes_last_week` decimal (12,3) comment '昨日收盘对比上周最高点涨跌幅0.1表示涨了10%,-0.1表示跌了10%',
	PRIMARY KEY (`id`),
	UNIQUE KEY `uidx_stock_basic_history` (`symbol`,`trade_date`)
); 
