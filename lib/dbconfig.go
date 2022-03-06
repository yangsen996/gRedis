package lib

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GORMDB *gorm.DB

func init() {
	GORMDB = gormDB()
}

func gormDB() *gorm.DB {
	dsn := ""
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	mysqlDB, _ := db.DB()

	mysqlDB.SetMaxIdleConns(5)
	mysqlDB.SetMaxOpenConns(10)
	return db

}
