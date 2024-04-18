package wxgo

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/skip2/go-qrcode"
)

func GenerateRandomTicket(length int) string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func GenerateQrCode(url string, savedir string, fname string) error {

	qrcode, err := qrcode.New(url, qrcode.Highest)
	if err != nil {
		return err
	}
	qrcode.DisableBorder = true
	//保存成文件
	savepath := fmt.Sprintf("%s/%s.png", savedir, fname)
	err = qrcode.WriteFile(256, savepath)
	return err
}
