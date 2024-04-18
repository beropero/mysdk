package wxgo

import "time"


type (
	// 微信服务器验证
	VerifyParams struct {
		Signature string
		Timestamp string
		Nonce     string
		Echostr   string
	}
	// access token
	AT struct {
		AccessToken string `json:"access_token"`
		ExpiresTime int    `json:"expires_in"`
		Time        time.Time
	}
	// 用户事件
	UserEvent struct {
		ToUserName   string `p:"ToUserName" xml:"ToUserName"`   // 开发者微信号
		FromUserName string `p:"FromUserName" xml:"FromUserName"` // 发送方账号（一个OpenID）
		CreateTime   int    `p:"CreateTime" xml:"CreateTime"`   // 消息创建时间（整型）
		MsgType      string `p:"MsgType" xml:"MsgType"`      // 消息类型，event,
		Event        string `p:"Event" xml:"Event"`        // 事件类型，subscribe（关注）, SCAN（已关注）
		EventKey     string `p:"EventKey" xml:"EventKey"`     // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
		Ticket       string `p:"Ticket" xml:"Ticket"`       // 二维码的ticket，可用来换取二维码图片
	}
	// 网页授权：access token
	UAT struct {
		AccessToken    string `json:"access_token"`
		ExpiresIn      int    `json:"expires_in"`
		RefreshToken   string `json:"refresh_token"`
		OpenId         string `json:"openid"`
		Scope          string `json:"scope"`
		IsSnapshotuser int    `json:"is_snapshotuser"`
		Unionid        string `json:"unionid"`
	}
	// 用户信息
	UserInfo struct {
		OpenId     string   `json:"openid"`
		NickName   string   `json:"nickname"`
		Sex        int      `json:"sex"`
		Province   string   `json:"province"`
		City       string   `json:"city"`
		Country    string   `json:"country"`
		Headimgurl string   `json:"headimgurl"`
		Privilege  []string `json:"privilege"`
		Unionid    string   `json:"unionid"`
	}
	
)

func NewAT() *AT {
	return &AT{
		AccessToken: "",
		ExpiresTime: 0,
		Time:        time.Now(),
	}
}
