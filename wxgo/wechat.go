package wxgo

import (
	"time"
)

type Wechat struct {
	Cfg      *WechatCfg
	LatestAT *AT
}
func NewWechat(cfg *WechatCfg) (*Wechat) {
	w := &Wechat{
		Cfg:      cfg,
		LatestAT: NewAT(),
	}
	err := w.RefreshAT()
	if err != nil {
		panic(err)
	}
	return w
}
// 刷新 LatestAT
func (w *Wechat) RefreshAT() error {
	// 先判断上次获取的是否超时
	duration := time.Since(w.LatestAT.Time)
	durationInSeconds := int(duration.Seconds())
	if durationInSeconds < (w.LatestAT.ExpiresTime - 600) {
		return nil
	}
	err := w.GetATReq()
	return err
}
// 获取 accesstoken
func (w *Wechat) GetAccessToken() (at string, err error) {
	w.RefreshAT()
	return w.LatestAT.AccessToken, nil
}
