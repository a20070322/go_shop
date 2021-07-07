package wechat

import (
	"context"
	"github.com/a20070322/shop-go/types/config"
	"github.com/a20070322/shop-go/global"

	"github.com/imroc/req"
	"time"
)

const ACCESS_TOKEN = "ACCESS_TOKEN"

type MiniChat struct {
	AppId  string
	secret string
}

func (m *MiniChat) Init(config *config.MinChat) {
	m.AppId = config.AppId
	m.secret = config.Secret
	//m.GetCatchToken()
	m.GetCatchToken()
}

// 获取缓存token
func (m *MiniChat) GetCatchToken() (string, error) {
	tokenCache, cacheErr := global.Rdb.Get(context.Background(), ACCESS_TOKEN).Result()
	if cacheErr == nil && tokenCache != "" {
		return tokenCache, nil
	}
	token, err := m.GetAccessToken()
	if err == nil {
		return token, nil
	}
	return "", err
}

// 获取小程序全局唯一后台接口调用凭据（access_token）。调用绝大多数后台接口时都需使用 access_token，开发者需要进行妥善保存。
func (m *MiniChat) GetAccessToken() (token string, err error) {
	var result ResGetAccessToken
	r, e := req.Get(GetUrl("GetAccessToken"), req.Param{
		"grant_type": "client_credential",
		"appid":      m.AppId,
		"secret":     m.secret,
	})
	if err := ResFn(r, e, &result); err != nil {
		return "", err
	}
	global.Rdb.Set(context.Background(), ACCESS_TOKEN, result.AccessToken, time.Second*7000)
	global.Logger.Info("获取access_token，有效时长7200")
	return result.AccessToken, nil
}

type ResCode2Session struct {
	CommonError
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
}

func (m *MiniChat) Code2Session(jsCode string) (result ResCode2Session, err error) {
	r, err := req.Get(GetUrl("Code2Session"), req.Param{
		"grant_type": "authorization_code",
		"appid":      m.AppId,
		"secret":     m.secret,
		"js_code":    jsCode,
	})
	if err := ResFn(r, err, &result); err != nil {
		return result, err
	}
	return result, nil
}
