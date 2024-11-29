package easemobimserversdk

import (
	"easemob-im-server-sdk/chatroom"
	"easemob-im-server-sdk/token"
)

type EasemobConfig struct {
	ApiUrl       string // 接口请求的地址
	OrgName      string // 环信企业名
	AppName      string // 环信应用名
	AppKey       string // 环信应用key
	ClientId     string
	ClientSecret string
}

type Easemob struct {
	Config   *EasemobConfig
	ChatRoom chatroom.IChatroom
	Token    token.IToken
}

// New  Easemob client
func New(config *EasemobConfig) *Easemob {
	return &Easemob{
		Config:   config,
		ChatRoom: chatroom.New(),
	}
}
