package database

import (
	"bluesdp/config"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func ReturnSession() *gorm.DB {

	var DBSession *gorm.DB
	//  this is sqlite connection
	db, _ := gorm.Open(sqlite.Open(config.Config("DATA_STORE_URI")), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(4)
	sqlDB.SetConnMaxLifetime(15 * time.Second)

	DBSession = db

	return DBSession

}
