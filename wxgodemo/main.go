package main

import (
	"mysdk/wxgodemo/router"

	"github.com/gin-gonic/gin"
)



func main() {
	r := gin.Default()
	router.RegisteredRoute(r)
	r.Static("/static","resource")
	r.Run(":80")
}