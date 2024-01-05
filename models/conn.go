package models

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var DSN = "root:password123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"

var DB1 *gorm.DB
var DSN1 = "xxx:xxxx@tcp(172.20.34.88:12325)/yonyou_cloud?charset=utf8mb4&parseTime=True&loc=Local"

func init() {
	var err error
	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               DSN, // data source name
		DefaultStringSize: 256, // default size for string fields

	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Printf("GORM init error: %v", err)
		return
	}
	setPool(DB)

	DB1, err = gorm.Open(mysql.New(mysql.Config{
		DSN:               DSN1, // data source name
		DefaultStringSize: 256,  // default size for string fields

	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Printf("GORM init error: %v", err)
		return
	}
	setPool(DB1)

}

func setPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Hour)

}
