package service

import (
	"accounting/repository"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type IUserService interface {
	Login(code string) (string, error)
}

func NewUserService() IUserService {
	return &UserService{}
}

type UserService struct {
	userRepo *repository.UserRepository
}

func (u UserService) Login(code string) (string, error) {
	body, err := getOpenidByCode(code)
	if err != nil {
		return "", err
	}

	var data map[string]string
	if err = json.Unmarshal(body, &data); err != nil {

	}
	user := u.userRepo.GetUserByOpenid(data["openid"])
	if user == nil {
		// TODO
		//u.userRepo.InsertUser(&datamodals.User{OpenId:data["openid"], })
	}

	return string(body[:]), nil
}

func getOpenidByCode(code string) ([]byte, error) {
	url := "https://api.weixin.qq.com/sns/jscode2session?appid=&secret="
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



