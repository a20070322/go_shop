package wechat

import (
	"fmt"
	"github.com/a20070322/shop-go/global"

	"github.com/imroc/req"
	"reflect"
	"strconv"
)

const  env = 1
var DevConfig = map[string]string{
	"GetAccessToken": "http://localhost:3000/cgi-bin/token",
	"Code2Session":"http://localhost:3000/sns/jscode2session",
}
var OnlineConfig = map[string]string{
	"GetAccessToken": "https://api.weixin.qq.com/cgi-bin/token",
	"Code2Session":"https://api.weixin.qq.com/sns/jscode2session",
}

func GetUrl(key string) string {
	if env == 0 {
		return DevConfig[key]
	}
	if env == 1 {
		return OnlineConfig[key]
	}
	return ""
}
//返回统一处理函数
func ResFn(res *req.Resp, err error, result interface{}) error {
	if err != nil {
		return err
	}
	res.ToJSON(&result)
	t := reflect.ValueOf(result).Elem()
	errCode := t.FieldByName("ErrCode")
	if errCode.Int() > 0 {
		errMsg := t.FieldByName("ErrMsg")
		errStr := fmt.Sprintf("code:%s,errmsg:%s", strconv.FormatInt(errCode.Int(), 10), errMsg.String())
		global.Logger.Error(errStr)
		return fmt.Errorf(errStr)
	}
	return nil
}
