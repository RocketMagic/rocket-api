package util

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB *gorm.DB

func init() {
	dbUrl := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", GetConfig("mysql.user"), GetConfig("mysql.password"), "tcp", GetConfig("mysql.host"), GetConfig("mysql.port"), "rocket", "utf8")
	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Printf("【数据库连接失败】：%s\n", err)
	}

	//defer db.Close()
	DB = db
}
