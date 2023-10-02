package database

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/goylold/lowcode/config"
	"xorm.io/xorm"
)

var engine *xorm.Engine
var engine2 *xorm.Engine

func init() {
	configBase := config.GetConConfig()
	var driverName string = configBase.Database.DriverName

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		configBase.Database.Username,
		configBase.Database.Pwd,
		configBase.Database.Host,
		configBase.Database.Port,
		configBase.Database.DatabaseName,
	)

	dataSourceInfoName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		configBase.Database.Username,
		configBase.Database.Pwd,
		configBase.Database.Host,
		configBase.Database.Port,
		configBase.Database.TableSchemaName,
	)

	var err error
	engine, err = xorm.NewEngine(driverName, dataSourceName)
	if err != nil {
		log.Panic("数据库加载失败!", err.Error())
	}
	err = engine.Ping()
	if err != nil {
		log.Panic("测试数据库连通性失败!", err.Error())
	}
	engine.ShowSQL(configBase.Database.ShowSQL)

	engine2, err = xorm.NewEngine(driverName, dataSourceInfoName)
	if err != nil {
		log.Panic("加载数据源2失败!", err.Error())
	}
	engine2.ShowSQL(configBase.Database.ShowSQL)
	err = engine2.Ping()
	if err != nil {
		log.Panic("测试数据源2连通性失败!", err.Error())
	}
}

func GetXOrmEngine() *xorm.Engine {
	return engine
}

func GetTableInfoEngine() *xorm.Engine {
	return engine2
}

func LoadDatabase() {
	log.Println("数据库连接加载中...")
}
