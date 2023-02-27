package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var (
	DB *gorm.DB
)

const (
	username = "root"
	password = "010110"
	host     = "127.0.0.1"
	dbName   = "bubble"
)

func InitMySQL() (err error) {

	dsn := fmt.Sprintf("%v:%v@(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, dbName)

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = DB.DB().Ping()
	return
}

func Close() {
	DB.Close()
}
