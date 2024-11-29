package chatroom

import (
	"context"
	"easemob-im-server-sdk/config"
	"easemob-im-server-sdk/request"
	"errors"
	"net/http"
)

// IChatroom 聊天室 提供的接口能力
// https://doc.easemob.com/document/server-side/chatroom_manage.html#创建聊天室
type IChatroom interface {
	Create(ctx context.Context, reqParam *CreateReq) (*CreateResp, error) // 创建聊天室
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
