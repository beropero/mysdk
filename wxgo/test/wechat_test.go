package wxgo_test

import (
	"testing"

	"github.com/beropero/mysdk/wxgo"
)


func Test_GetAccessToken(t *testing.T) {
	cfg := &wxgo.WechatCfg{
		Appid:     "wx546099fd0bd1abf6",
		Appsecret: "a5b71fdd20f02e312fc81591ce7c6af8",
		Token:     "beropero",
	}
	w := wxgo.NewWechat(cfg)
	AT1, err := w.GetAccessToken()
	if err != nil {
		t.Log(err)
	}
	t.Log(AT1)
	// 手动超时
	w.LatestAT.ExpiresTime = 0
	AT2, _ := w.GetAccessToken()
	t.Log(AT2)
	AT3, _ := w.GetAccessToken()
	t.Log(AT3)
	if AT1 == "" {
		t.Fatal("获取失败")
	} else if AT1 == AT2 {
		t.Fatal("超时获取失败")
	} else if AT2 != AT3 {
		t.Fatal("未超时重复获取无效")
	}
}