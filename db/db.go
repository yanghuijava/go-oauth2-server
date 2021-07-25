package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
)

var db *gorm.DB

func Close() {
	if db != nil {
		logrus.Info("关闭数据连接")
		db.Close()
	}
}

func GetDb() *gorm.DB {
	return db
}

func InitDb() {
	d, err := gorm.Open("mysql", "root:123456@(10.100.0.116:30113)/oauth2-server?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	d.DB().SetMaxOpenConns(10)
	d.DB().SetMaxIdleConns(5)
	db = d
	db.Debug()
	db.SingularTable(true)
	db.LogMode(true)
	logrus.Info("数据库连接成功")
}
