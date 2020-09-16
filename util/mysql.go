package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const DSN = "root:root@tcp(106.14.5.37)/AccountingSystem?charset=utf8&parseTime&loc=Local"

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