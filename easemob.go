package easemobimserversdk

import (
	"errors"
	"fmt"

	"github.com/jjeejj/easemob-im-server-sdk/chatroom"
	"github.com/jjeejj/easemob-im-server-sdk/config"
	"github.com/jjeejj/easemob-im-server-sdk/request"
	"github.com/jjeejj/easemob-im-server-sdk/token"
	"github.com/jjeejj/easemob-im-server-sdk/user"
)

type Easemob struct {
	Config     *config.EasemobConfig
	ChatRoom   chatroom.IChatroom
	Token      token.IToken
	User       user.IUser
	httpClient *request.HttpClient // 发送请求的客户端
}

// New  Easemob client
func New(config *config.EasemobConfig) (*Easemob, error) {
	if config.ApiUrl == "" || config.AppKey == "" || config.ClientId == "" || config.ClientSecret == "" || config.OrgName == "" || config.AppName == "" {
		return nil, errors.New("config error")
	}
	httpClient := request.New(fmt.Sprintf("%s/%s/%s", config.ApiUrl, config.OrgName, config.AppName))
	return &Easemob{
		Config:     config,
		httpClient: httpClient,
		ChatRoom:   chatroom.New(config, httpClient),
		Token:      token.New(config, httpClient),
		User:       user.New(config, httpClient),
	}, nil
}
