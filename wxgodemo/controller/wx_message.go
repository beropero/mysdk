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

func WxMessage(ctx *gin.Context) {
	vp := &wxgo.VerifyParams{
		Signature: ctx.Query("signature"),
		Echostr:   ctx.Query("echostr"),
		Timestamp: ctx.Query("timestamp"),
		Nonce:     ctx.Query("nonce"),
	}
	if flag, _ := wechat.Wx.VerifySignature(*vp); flag{
		log.Println(vp)
		ctx.String(http.StatusOK, vp.Echostr)
	}
	uEvent := &wxgo.UserEvent{}
	ctx.ShouldBindXML(&uEvent)
	fmt.Println(uEvent)
	if uEvent.Ticket != "" && (uEvent.Event == "SCAN"||uEvent.Event=="subscribe"){	
		openid  := uEvent.FromUserName
		log.Printf("ticket:%s, openid:%s",uEvent.Ticket, openid)
		// 将ticket和openid存入redis
		err := redis.RedisClient.Set(uEvent.Ticket, openid, 60*time.Second).Err()
		if err != nil {
			log.Println(err.Error())
		}
	}
}
