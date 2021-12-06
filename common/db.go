package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync"
)

var db *gorm.DB
var once sync.Once

func StartDatabase(username string, password string, dbname string, host string, port string) error {
	var err error
	var tdb *gorm.DB

	once.Do(func() {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password,
			host, port, dbname)
		tdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		db = tdb
	})
	if err != nil {
		return err
	}
	return nil
}

func GetDb() *gorm.DB {
	return db
}
