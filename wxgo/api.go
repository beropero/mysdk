package wxgo

type (

	// QR code 请求体
	QRCodeReq struct {
		ExpireSeconds int        `json:"expire_seconds"`
		ActionName    string     `json:"action_name"`
		ActionInfo    ActionInfo `json:"action_info"`
	}
	ActionInfo struct {
		Scene `json:"scene"`
	}
	Scene struct {
		SceneId int `json:"scene_id"`
	}
	// QR code 响应体
	QRCodeRes struct {
		Ticket        string `json:"ticket"`
		ExpireSeconds int    `json:"expire_seconds"`
		Url           string `json:"url"`
	}
)
