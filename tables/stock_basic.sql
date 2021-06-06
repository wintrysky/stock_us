/*
SQLyog 企业版 - MySQL GUI v8.12,3 
MySQL - 5.7.18-20170830-log 
*********************************************************************
*/
/*!40101 SET NAMES utf8 */;

create table `stock_basic` (
	`id` bigint NOT NULL AUTO_INCREMENT,
	`symbol` varchar (100) comment '代码',
	`name` varchar (200) comment '名称',
	`industry` varchar (100) comment '所属行业',
	`is_china` varchar (1) comment '中概股',
	`is_hot` varchar (1) comment '热门',
	`is_etf` varchar (1) comment '期权',
	`company_tags` varchar (20) comment '护城河,优质股,潜力股,十倍股',
	`description` text comment '描述',
	`enabled` varchar (1) default 'Y' comment '是否可用',
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
	`trade_date` datetime comment '股票交易日',
	`calcute_date` datetime comment '计算周涨跌幅完成时间,如果无效时间,表示不可交易',
	`compare_last_week` decimal (12,3) comment '对比上周涨跌幅0.1表示涨了10%,-0.1表示跌了10%',
	`compare_this_week` decimal (12,3) comment '对比周五和周一涨跌幅0.1表示涨了10%,-0.1表示跌了10%',
	`create_time` datetime ,
	`con_id` bigint comment 'IB股票标记值',
	`sec_type` varchar (20) comment 'STK/OPT',
	`prim_exchange` varchar (150) comment 'NASDAQ.NMS/NYSE',
	`exchange` varchar (20) comment 'ISLAND/SMART',
	`currency` varchar (20) comment 'USD',
	PRIMARY KEY (`id`),
	UNIQUE KEY `uidx_stock_basic_symbol` (`symbol`)
); 
