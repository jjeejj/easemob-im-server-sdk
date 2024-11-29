package easemobimserversdk

import (
	"easemob-im-server-sdk/chatroom"
	"easemob-im-server-sdk/config"
	"easemob-im-server-sdk/request"
	"easemob-im-server-sdk/token"
	"fmt"
)

type Easemob struct {
	Config     *config.EasemobConfig
	ChatRoom   chatroom.IChatroom
	Token      token.IToken
	httpClient *request.HttpClient // 发送请求的客户端
}

// New  Easemob client
func New(config *config.EasemobConfig) *Easemob {
	if config.ApiUrl == "" || config.AppKey == "" || config.ClientId == "" || config.ClientSecret == "" || config.OrgName == "" || config.AppName == "" {
		panic("config error")
	}
	httpClient := request.New(fmt.Sprintf("%s/%s/%s", config.ApiUrl, config.OrgName, config.AppName))
	return &Easemob{
		Config:     config,
		httpClient: httpClient,
		ChatRoom:   chatroom.New(config, httpClient),
		Token:      token.New(config, httpClient),
	}
}
