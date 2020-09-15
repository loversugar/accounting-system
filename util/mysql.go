package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DSN = "root:root@tcp()/accounting-system?charset=utf8&parseTime&loc=Local"

func Connection() (db *gorm.DB, err error) {
	return gorm.Open(mysql.New(mysql.Config{
		DSN: DSN,
	}), &gorm.Config{})
}