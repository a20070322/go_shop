package initialize

import (
	"github.com/a20070322/shop-go/global"
	"github.com/a20070322/shop-go/global/ws"
	"github.com/a20070322/shop-go/pkg/wechat"
)

func MiniChat() {
	m := &wechat.MiniChat{}
	m.Init(global.AppSetting.MinChat)
	ws.MiniChat = m
}
