package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

func login(context *gin.Context)  {
	code := context.Query("code")
	logrus.Info(code)

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
}

func register(context *gin.Context)  {

}
