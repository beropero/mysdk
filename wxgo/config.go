package wxgo

import "errors"

type WechatCfg struct {
	Token       string
	Appid       string
	Appsecret   string
	ExpiresTime int
}

func (cfg *WechatCfg) GetAppid() (string, error) {
	if cfg.Appid == "" {
		return "", errors.New("get config error: appid undefined")
	}
	return cfg.Appid, nil
}

func (cfg *WechatCfg) GetAppsecret() (string, error) {
	if cfg.Appsecret == "" {
		return "", errors.New("get config error: appsecret undefined")
	}
	return cfg.Appsecret, nil
}

func (cfg *WechatCfg) GetToken() (string, error) {
	if cfg.Token == "" {
		return "", errors.New("get config error: appsecret undefined")
	}
	return cfg.Token, nil
}

func (cfg *WechatCfg) GetExpiresTime() int {
	if cfg.ExpiresTime == 0 {
		cfg.ExpiresTime = 60
	}
	return cfg.ExpiresTime
}
