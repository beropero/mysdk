package router

import (
	"mysdk/wxgodemo/controller"

	"github.com/gin-gonic/gin"
)

func RegisteredRoute(r *gin.Engine){
	wxgroup := r.Group("/wechat")
	{
		wxgroup.Any("/message",controller.WxMessage)
		wxgroup.GET("/getloginqr",controller.GetGZQrUrl)
		wxgroup.GET("/getauthqr",controller.GetAuthQrUrl)
		wxgroup.GET("/accessusercode",controller.AccessUserCode)
		wxgroup.GET("/checklogin",controller.CheckLogin)
	}
}