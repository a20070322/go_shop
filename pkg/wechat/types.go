package wechat


type CommonRep struct{
	ErrCode int64  `json:"errcode"`
}

type CommonError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
type ResGetAccessToken struct {
	CommonError
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}