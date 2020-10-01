package datamodals

import "time"

type User struct {
	ID int
	OpenId string `gorm:"primaryKey;column:openid"`
	UserName string `gorm:"column:username"`
	CreateTime time.Time
}
