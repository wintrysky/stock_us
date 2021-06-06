package csv

import (
	"bufio"
	"encoding/csv"
	"github.com/wintrysky/stock_us/template"
	"github.com/wintrysky/stock_us/utils"
	"github.com/wintrysky/stock_us/xerr"
	"os"
	"strings"
)

type FileCsvService struct {
	headerMap map[int]string           // 列位置-列名
	dataMap   []map[string]interface{} // english column name -- value
}

// Load CSV file mapper.MapperMap(data, &info)
func (c *FileCsvService)ImportCSV(filePath string) (data []map[string]interface{},err error) {
	defer xerr.HandleErr(&err)
	// get all rows from csv
	records := c.readData(filePath)
	// set head position/name to map
	c.setHeaderMap(records)
	// convert csv rows to []map[col]value
	// caller should convert them to database models from outside
	c.convertToMap(records)

	return c.dataMap,nil
}

func (c *FileCsvService)readData(filePath string)(records [][]string){
	file, err := os.Open(filePath)
	xerr.ThrowError(err)

	rr := bufio.NewReader(file)
	tsv := csv.NewReader(rr)
	//tsv.Comma = '\t'
	tsv.Comment = '#'
	tsv.LazyQuotes = true
	//tsv.TrailingComma = true     // retain rather than remove empty slots
	tsv.TrimLeadingSpace = false // retain rather than remove empty slots

	records, err = tsv.ReadAll()
	xerr.ThrowError(err)

	return records
}

func (c *FileCsvService)setHeaderMap(records [][]string){
	c.headerMap = make(map[int]string) // csv列位置--中文标题

	for _, rec := range records {
		for idx, name := range rec {
			name1 := strings.ReplaceAll(utils.TrimSpace(name), "\"", "")
			name1 = utils.TrimSpace(name1)
			c.headerMap[idx] = name1
		}
		break
	}
}

func (c *FileCsvService)convertToMap(records [][]string){
	for idx, rec := range records {
		if idx == 0 {
			continue
		}
		c.convertRow(rec)
	}
}

func (c *FileCsvService)convertRow(record []string){
	mm := make(map[string]interface{})
	for idx, value := range record {
		if columnName, ok := c.headerMap[idx]; ok {
			// get english name from chinese
			key := c.getEnglishColumnName(columnName)
			if key != "" {
				c.setColumnValue(mm,key, strings.TrimSpace(value))
			}
		}
	}

	c.dataMap = append(c.dataMap,mm)
}

func (c *FileCsvService)getEnglishColumnName(key string) string {
	if v, ok := template.FutuColumnMap[key]; ok {
		return v
	}

	return ""
}

// key = engColumnName,datetime
func (c *FileCsvService)setColumnValue(mm map[string]interface{},key string, value string) {
	arr := strings.Split(key,",")
	col := arr[0]
	columnType := arr[1] // string,stringToInt,datetime

	var v interface{}
	if  columnType == "string" {
		v = toString(value)
	}
	if  columnType == "float" {
		v = toFloat(value)
	}
	if  columnType == "int" {
		v = toInt(value)
	}
	if  columnType == "datetime" {
		v = toDatetime(value)
	}
	if  columnType == "percent" {
		v = toPercent(value)
	}
	if  columnType == "money" {
		v = toMoney(value)
	}

	mm[col] = v
}