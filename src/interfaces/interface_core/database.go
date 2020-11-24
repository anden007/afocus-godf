package interface_core

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IDataBase interface {
	GetDB() *gorm.DB
}
