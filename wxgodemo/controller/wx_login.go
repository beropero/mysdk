package controller

import (
	"fmt"
	"log"
	"mysdk/wxgodemo/redis"
	"mysdk/wxgodemo/wechat"
	"net/http"
	"time"
	"wxgo"

	"github.com/gin-gonic/gin"
)

// 登陆轮训接口
func CheckLogin(ctx *gin.Context) {
	ticket := ctx.Query("ticket")
	openid, err := redis.RedisClient.Get(ticket).Result()
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"login": false,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"login":  true,
			"openid": openid,
		})
	}
}

// 获取公众号登陆二维码
func GetGZQrUrl(ctx *gin.Context) {
	ticket, _ := wechat.Wx.GetQRTicketReq("QR_STR_SCENE", 123)
	qrUrl := wechat.Wx.GetQrImageUrl(ticket)
	ctx.JSON(http.StatusOK, gin.H{
		"ticket": ticket,
		"qrUrl":  qrUrl,
	})
}

// 获取网页授权登陆二维码
func GetAuthQrUrl(ctx *gin.Context) {
	// 生成ticket
	ticket := wxgo.GenerateRandomTicket(20)
	// 生成授权地址
	redirect_url := "http://39.101.78.10/wechat/accessusercode" // 微信授权后重定向地址,用于接收用户code
	scope := "snsapi_base"                                  //授权权限
	oauthUrl := wechat.Wx.GetOauth2CodeUrl(redirect_url, scope, ticket)
	// 将授权地址生成QR码
	savePath := "./resource/image"
	err := wxgo.GenerateQrCode(oauthUrl, savePath, ticket)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	qrUrl := fmt.Sprintf("http://39.101.78.10/static/image/%s.png", ticket)
	ctx.JSON(http.StatusOK, gin.H{
		"ticket": ticket,
		"qrUrl":  qrUrl,
	})
}

// 网页授权接收code
func AccessUserCode(ctx *gin.Context) {
	code := ctx.Query("code")
	ticke := ctx.Query("state")
	// 用code获取用户access token
	uat, _ := wechat.Wx.GetUserATReq(code)
	// 用access token获取用户信息
	uInfo, _ := wechat.Wx.GetUserInfoReq(uat)
	log.Printf(uInfo.NickName, uInfo.Headimgurl, uInfo.City)
	// 将ticket和openid放入redis
	redis.RedisClient.Set(ticke, uInfo.OpenId, 60*time.Second)
	// 重定向到成功/失败页面
	ctx.Redirect(http.StatusTemporaryRedirect, "http://39.101.78.10/static/html/loginsucceed.html")
}
