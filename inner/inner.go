package inner

import (
	"os"
	"time"
	"errors"
	"github.com/luaxlou/wx-qyapi-sdk/lib/httpclient"
)

const BASE_URL = "https://qyapi.weixin.qq.com/cgi-bin"

var (
	corpid     = os.Getenv("CORP_ID")
	corpsecret = os.Getenv("CORP_SECRET")
)

type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn   int       `json:"expires_in"`
	CreateTime  time.Time `json:"create_time"`
}

func (t *Token) IsExpired() bool {

	if time.Now().Unix()-t.CreateTime.Unix() > int64(t.ExpiresIn) {
		return true

	}
	return false
}

var currToken *Token

type GetTokenRes struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func getToken() (string, error) {

	if currToken == nil || currToken.AccessToken == "" || currToken.IsExpired() {

		var res *GetTokenRes

		url := BASE_URL + "/gettoken"

		if corpid == "" || corpsecret == "" {
			return "", errors.New("Please set CORP_ID CORP_SECRET on OS ENV.")
		}

		err := httpclient.GetJSON(url, map[string]string{"corpid": corpid, "corpsecret": corpsecret}, &res)

		if err != nil {
			return "", err

		}

		if res.Errcode != 0 {
			return "", errors.New(res.Errmsg)
		}

		currToken = &Token{AccessToken: res.AccessToken, ExpiresIn: res.ExpiresIn, CreateTime: time.Now()}

	}

	return currToken.AccessToken, nil

}

func get(uri string, params map[string]string, res interface{}) error {

	accessToken, err := getToken()
	if err != nil {
		return err
	}

	params["access_token"] = accessToken

	url := BASE_URL + uri + ""
	return httpclient.GetJSON(url, params, res)

}

func post(uri string, params map[string]string, res interface{}) error {

	accessToken, err := getToken()
	if err != nil {
		return err
	}

	url := BASE_URL + uri + "?access_token=" + accessToken
	return httpclient.PostJSON(url, params, res)

}

func postBody(uri string, req interface{}, res interface{}) error {

	accessToken, err := getToken()


	if err != nil {
		return err
	}

	url := BASE_URL + uri + "?access_token=" + accessToken
	return httpclient.PostBody(url, req, res)

}
