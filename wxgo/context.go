package wxgo

type Param interface{}

type WeContext struct {
	Ticket string
	Params map[string]Param
	QrUrl string
	UserInfo UserInfo
}