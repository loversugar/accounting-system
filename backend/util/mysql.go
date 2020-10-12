package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const DSN = "root:root@tcp(127.0.0.1)/AccountingSystem?charset=utf8&parseTime=true&loc=Local"

func Connection() (db *gorm.DB, err error) {
	return gorm.Open(mysql.New(mysql.Config{
		DSN: DSN,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
	})
}