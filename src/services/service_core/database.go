package service_core

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/anden007/afocus-godf/src/generate"
	"github.com/anden007/afocus-godf/src/lib"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DataBase struct {
	db *gorm.DB
}

func NewDataBase() *DataBase {
	instance := new(DataBase)
	loadTime := time.Now()
	var myLogger logger.Interface
	if lib.IS_DEV_MODE {
		myLogger = logger.Default.LogMode(logger.Info)
	} else {
		myLogger = logger.Default.LogMode(logger.Silent)
	}
	dsn := os.Getenv("mysql-server")
	if mySqlDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                                   myLogger,
		DisableForeignKeyConstraintWhenMigrating: true,
	}); err == nil {
		mySqlDB.Set("gorm:table_options", "ENGINE="+os.Getenv("mysql-engine"))
		if db, err := mySqlDB.DB(); err == nil {
			db.SetMaxIdleConns(10)
			db.SetMaxOpenConns(100)
			db.SetConnMaxLifetime(time.Minute * 30)
		}
		if lib.IS_DEV_MODE && strings.EqualFold(os.Getenv("auto-migrate"), "true") {
			generate.AutoMigrate(mySqlDB)
		}
		instance.db = mySqlDB
	} else {
		fmt.Println("Connect to MySQL error", err)
		return nil
	}
	if lib.IS_DEV_MODE {
		fmt.Println("> Service: DataBase loaded.", time.Now().Sub(loadTime))
	}
	return instance
}

func (m *DataBase) GetDB() *gorm.DB {
	return m.db
}
