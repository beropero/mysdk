package wechat

import "wxgo"

var (
	cfg = &wxgo.WechatCfg{
		Appid:     "your_appid",
		Appsecret: "your_appsecret",
		Token:     "your_token",
	}
	Wx = wxgo.NewWechat(cfg)
)