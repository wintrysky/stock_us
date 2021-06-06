package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"github.com/wintrysky/orm"
	"github.com/wintrysky/stock_us/global"
	"github.com/wintrysky/stock_us/router"
	"github.com/wintrysky/stock_us/template"
	"github.com/wintrysky/stock_us/utils"
	"net/http"
)

func init() {
	utils.InitLog("./log/", "service", "gb18030")
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func startServer() {
	// Creates a router without any middleware by default
	r := gin.New()

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	r.Use(cors())

	router.UseRouters(r)

	var port int
	global.ReadKey("host.port",&port)
	host := fmt.Sprintf(":%v", port)

	r.Run(host)
}

func initORM() {
	config := make(map[string]interface{})
	global.ReadKey("db", &config)

	dbSettings := orm.DBSettings{}
	dbSettings.DBName = cast.ToString(config["dbname"])
	dbSettings.Password = cast.ToString(config["password"])
	dbSettings.Port = cast.ToInt(config["port"])
	dbSettings.User = cast.ToString(config["user"])
	dbSettings.Host = cast.ToString(config["host"])
	err := orm.InitDB(dbSettings)
	if err != nil {
		log.Error(err)
	}
}

func initTemplate(){
	template.InitFutuColumnMap()
}

func main(){
	initORM()
	initTemplate()
	startServer()
}
