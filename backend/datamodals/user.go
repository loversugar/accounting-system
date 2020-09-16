package datamodals

import "time"

type User struct {
	OpenId string `gorm:"primaryKey;column:openId"`
	UserName string `gorm:"column:username"`
	CreateTime time.Time
}
