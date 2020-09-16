package repository

import (
	"accounting/datamodals"
	"testing"
	"time"
)

func Test_InsertUser(t *testing.T) {
	userRepository := NewUserRepository()
	if err := userRepository.InsertUser(&datamodals.User{OpenId:"test3", UserName:"zs", CreateTime:time.Now()});
		err != nil{
		t.Error("插入数据失败")
		return
	}
	t.Log("插入成功")
}