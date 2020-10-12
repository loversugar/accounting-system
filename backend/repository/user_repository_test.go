package repository

import (
	"testing"
)

func Test_InsertUser(t *testing.T) {
	userRepo := NewUserRepository()
	//user := userRepo.GetUserByOpenid("ovnAO5P9e9oc-dHm37Lpd6pW1BlY")
	//t.Log(user.CreateTime)

	user := userRepo.GetUserByOpenid("test")
	t.Log(user)
}