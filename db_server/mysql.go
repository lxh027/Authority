package db_server

import (
	"Authority/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)


var MySqlDb *gorm.DB
var MySqlError error

func init()  {
	dbConfig := config.GetDbConfig()

	// set db dsn
	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s",
			dbConfig["username"],
			dbConfig["password"],
			dbConfig["hostname"],
			dbConfig["port"],
			dbConfig["database"],
			dbConfig["charset"],
			dbConfig["parseTime"],
		)

	// open connection
	MySqlDb, MySqlError = gorm.Open("mysql", dbDSN)

	// 禁用默认复数表名
	MySqlDb.SingularTable(true)

	if MySqlError != nil {
		panic("database open error! " + MySqlError.Error())
	}


}




