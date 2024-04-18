package wxgo

import (
	"fmt"
	"net/url"
)

type Url struct {
	ATUrl            string // access token 获取地址
	TicketUrl        string // ticket 获取地址
	QRImgUrl         string // qr code 图片地址
	UserInfoUrl      string // 网页授权：拉取用户信息
	Oauth2CodeUrl    string // 网页授权：Code
	UserATUrl        string // 网页授权：用户 access token
	RefreshUserATUrl string // 网页授权：刷新用户 access token
}

var (
	ReqUrl = &Url{
		ATUrl:            "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		TicketUrl:        "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token=%s",
		QRImgUrl:         "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s",
		UserInfoUrl:      "https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s&lang=zh_CN",
		Oauth2CodeUrl:    "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=%s&state=%s#wechat_redirect",
		UserATUrl:        "https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
		RefreshUserATUrl: "https://api.weixin.qq.com/sns/oauth2/refresh_token?appid=%s&grant_type=refresh_token&refresh_token=%s",
	}
)

func (w *Wechat) GetOauth2CodeUrl(redirectUrl string, scope string, state string) string {
	encodeUrl := url.QueryEscape(redirectUrl)
	url := fmt.Sprintf(ReqUrl.Oauth2CodeUrl, w.Cfg.Appid, encodeUrl, scope, state)
	return url
}

func (w *Wechat) GetUserATUrl(code string) string {
	url := fmt.Sprintf(ReqUrl.UserATUrl, w.Cfg.Appid, w.Cfg.Appsecret, code)
	return url
}

func (w *Wechat) GetUserInfoUrl(at UAT) string {
	url := fmt.Sprintf(ReqUrl.UserInfoUrl, at.AccessToken, at.OpenId)
	return url
}

func (w *Wechat) GetQrImageUrl(ticket string) string {
	url := fmt.Sprintf(ReqUrl.QRImgUrl,ticket)
	return url
}