package wxgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func (w *Wechat) GetQRTicketReq(codetype string, sceneId int) (string, error) {
	// 获取 at
	at, err := w.GetAccessToken()
	if err != nil {
		return "", err
	}
	// 拼接请求地址
	url := fmt.Sprintf(ReqUrl.TicketUrl, at)
	// 构造请求数据
	data := &QRCodeReq{
		ExpireSeconds: w.Cfg.GetExpiresTime(),
		ActionName:    codetype, // QR码 类型
		ActionInfo: ActionInfo{
			Scene: Scene{
				SceneId: sceneId,
			},
		},
	}
	// 发送 post 请求获取响应
	client := &http.Client{}
	Jsondata, err := json.Marshal(&data)
	if err != nil {
		return "", err
	}
	reader := bytes.NewReader(Jsondata)
	req, err := http.NewRequest("POST", url, reader)
	if err != nil {
		return "", err
	}
	// 设置请求头 json格式
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var respData = QRCodeRes{}
	// 解析字符串
	err = json.Unmarshal(body, &respData)
	if err != nil {
		return "", errors.New("json unmarsha fail")
	}
	return respData.Ticket, nil
}
// 获取用户 access token
func (w *Wechat) GetUserATReq(code string) (UAT, error) {
	url := w.GetUserATUrl(code)
	// 发送 GET 请求获取响应
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return UAT{}, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return UAT{}, err
	}
	defer resp.Body.Close()
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return UAT{}, err
	}

	uat := UAT{}
	json.Unmarshal(body, &uat)
	return uat, nil
}

func (w *Wechat) GetUserInfoReq(at UAT) (UserInfo, error) {
	url := w.GetUserInfoUrl(at)
	// 发送 GET 请求获取响应
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return UserInfo{}, err
	}
	// 设置请求头 json格式
	resp, err := client.Do(req)
	if err != nil {
		return UserInfo{}, err
	}
	defer resp.Body.Close()
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserInfo{}, err
	}
	uInfo := UserInfo{}
	json.Unmarshal(body, &uInfo)
	return uInfo, nil
}
// 获取普通access token
func (w *Wechat) GetATReq() error {
	// 获取配置信息
	appid, err := w.Cfg.GetAppid()
	if err != nil {
		return err
	}
	appsecret, err := w.Cfg.GetAppsecret()
	if err != nil {
		return err
	}
	// 构造请求地址
	url := fmt.Sprintf(ReqUrl.ATUrl, appid, appsecret)
	// 发送 GET 请求获取响应
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 读取响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	// 如果响应结果包含错误错误码，返回错误信息
	if strings.Contains(string(body), "errcode") {
		return fmt.Errorf("wechat response error: %s", string(body))
	}
	// 解析字符串
	err = json.Unmarshal(body, &w.LatestAT)
	if err != nil {
		return errors.New("json Unmarshal fail")
	}
	// 设置成功获取时间
	w.LatestAT.Time = time.Now()
	return nil
}
