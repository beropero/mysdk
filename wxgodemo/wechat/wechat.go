package wechat

import "wxgo"

var (
	cfg = &wxgo.WechatCfg{
		Appid:     "wx546099fd0bd1abf6",
		Appsecret: "a5b71fdd20f02e312fc81591ce7c6af8",
		Token:     "beropero",
	}
	Wx = wxgo.NewWechat(cfg)
)