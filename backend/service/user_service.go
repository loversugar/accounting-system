package service

import (
	"accounting/datamodals"
	"accounting/repository"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type IUserService interface {
	Login(code, nickName string) (string, error)
}

func NewUserService() IUserService {
	return &UserService{userRepo:repository.NewUserRepository()}
}

type UserService struct {
	userRepo repository.IUserRepository
}

func (u UserService) Login(code, nickName string) (string, error) {
	body, err := getOpenidByCode(code)
	if err != nil {
		return "", err
	}

	var data map[string]string
	if err = json.Unmarshal(body, &data); err != nil {
		return "", err
	}
	user := u.userRepo.GetUserByOpenid(data["openid"])
	if user == nil {
		timeNow := time.Now()
		user, err = u.userRepo.InsertUser(
			&datamodals.User{OpenId:data["openid"], UserName:nickName, CreateTime: timeNow})
		if err != nil {
			return "", err
		}
	}
	data["userId"] = strconv.Itoa(user.ID)
	dataType, _ := json.Marshal(data)
	dataString := string(dataType)

	return dataString, nil
}

func getOpenidByCode(code string) ([]byte, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=wx10ce95ecd4642a6b&secret=7bd9ffe3049cd21632b3fdec8bd59e45"
	url += "&js_code="+code+"&grant_type=authorization_code"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("content-type", "application/json")
	if err != nil {
		logrus.Error(err)
	}
	resp, err := client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	logrus.Info(string(body[:]))
	return body, err
}



