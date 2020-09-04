package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"golangPro/golang-mallapi/config"
	"log"
)

type BaseModel struct {
	DbConnect *gorm.DB
}

func (baseModel BaseModel) GetDbConnect(mysqlDbKey string) *gorm.DB {
	var err error
	if mysqlDbKey == "" {
		mysqlDbKey = "default"
	}
	mysqlConfig := config.GetMysqlConf(mysqlDbKey)
	baseModel.DbConnect, err = gorm.Open(config.DriverName, mysqlConfig["user"]+":"+mysqlConfig["password"]+"@tcp("+mysqlConfig["host"]+":"+mysqlConfig["port"]+")/"+mysqlConfig["dbname"]+"?charset=utf8&parseTime=True&loc=Local")
	//baseModel.DbConnect.LogMode(config.DbLoG)
	//defer baseModel.dbConnect.Close()
	if err != nil {
		// 打日志
		log.Println("数据库连接错误----", err)
		return nil
	}
	return baseModel.DbConnect
}

func (baseModel BaseModel) Close() {
	if baseModel.DbConnect != nil {
		baseModel.DbConnect.Close()
	}
}
