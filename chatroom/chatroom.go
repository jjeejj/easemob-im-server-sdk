package chatroom

import (
	"context"
	"errors"
	"net/http"

	"github.com/jjeejj/easemob-im-server-sdk/config"
	"github.com/jjeejj/easemob-im-server-sdk/request"
)

// IChatroom 聊天室 提供的接口能力

type IChatroom interface {
	// https://doc.easemob.com/document/server-side/chatroom_manage.html#创建聊天室
	Create(ctx context.Context, reqParam *CreateReq) (*CreateResp, error) // 创建聊天室
	// https://doc.easemob.com/document/server-side/message_chatroom.html#http-请求-3
	SendMessage(ctx context.Context, reqParam *SendMessageReq) (*SendMessageResp, error) // 发送消息
}

type Chatroom struct {
	httpClient *request.HttpClient // 发送请求的客户端
	config     *config.EasemobConfig
}

func New(config *config.EasemobConfig, httpClient *request.HttpClient) IChatroom {
	return &Chatroom{
		httpClient: httpClient,
		config:     config,
	}
}

func (c *Chatroom) Create(ctx context.Context, reqParam *CreateReq) (*CreateResp, error) {
	var respData *CreateResp = &CreateResp{}
	restResponse, err := c.httpClient.Post(ctx, ApiCreateChatroomPath, map[string]any{
		"name":        reqParam.Name,
		"description": reqParam.Description,
		"maxusers":    reqParam.MaxUsers,
		"owner":       reqParam.Owner,
		"members":     reqParam.Members,
		"custom":      reqParam.Custom,
	}, map[string]string{
		"Authorization": "Bearer " + reqParam.AppToken,
	}, respData)
	if restResponse.StatusCode() != http.StatusOK {
		return nil, errors.New(restResponse.String())
	}
	return respData, err
}

func (c *Chatroom) SendMessage(ctx context.Context, reqParam *SendMessageReq) (*SendMessageResp, error) {
	var respData *SendMessageResp = &SendMessageResp{}
	restResponse, err := c.httpClient.Post(ctx, ApiSendMessagePath, reqParam, map[string]string{
		"Authorization": "Bearer " + reqParam.AppToken,
	}, respData)
	if restResponse.StatusCode() != http.StatusOK {
		return nil, errors.New(restResponse.String())
	}
	return respData, err
}
