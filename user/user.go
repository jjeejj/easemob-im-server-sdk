package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/jjeejj/easemob-im-server-sdk/config"
	"github.com/jjeejj/easemob-im-server-sdk/request"
)

// https://doc.easemob.com/document/server-side/account_system.html#批量授权注册用户
type IUser interface {
	Registry(ctx context.Context, reqParam *RegistryReq) (*RegistryResp, error)                // 注册一个用户
	BatchRegistry(ctx context.Context, reqParam *BatchRegistryReq) (*BatchRegistryResp, error) // 批量注册用户，单次请求最多可注册 60 个用户 ID。
}

type User struct {
	config     *config.EasemobConfig
	httpClient *request.HttpClient // 发送请求的客户端
}

func New(config *config.EasemobConfig, httpClient *request.HttpClient) IUser {
	return &User{
		httpClient: httpClient,
		config:     config,
	}
}

func (u *User) Registry(ctx context.Context, reqParam *RegistryReq) (*RegistryResp, error) {
	var respData *RegistryResp = &RegistryResp{}
	restResponse, err := u.httpClient.Post(ctx, ApiRegistryUserPath, map[string]any{
		"username": reqParam.User.Username,
		"password": reqParam.User.Password,
	}, map[string]string{
		"Authorization": "Bearer " + reqParam.AppToken,
	}, respData)
	if restResponse.StatusCode() != http.StatusOK {
		return nil, errors.New(restResponse.String())
	}
	return respData, err
}

func (u *User) BatchRegistry(ctx context.Context, reqParam *BatchRegistryReq) (*BatchRegistryResp, error) {
	// 用户信息必传
	if len(reqParam.Users) == 0 {
		return nil, errors.New("users is empty")
	}
	var respData *BatchRegistryResp = &BatchRegistryResp{}
	batchRegistryUserReqParam := make([]map[string]any, 0)
	for _, user := range reqParam.Users {
		batchRegistryUserReqParam = append(batchRegistryUserReqParam, map[string]any{
			"username": user.Username,
			"password": user.Password,
		})
	}
	restResponse, err := u.httpClient.Post(ctx, ApiBatchRegistryUserPath, batchRegistryUserReqParam, map[string]string{
		"Authorization": "Bearer " + reqParam.AppToken,
	}, respData)
	if restResponse.StatusCode() != http.StatusOK {
		return nil, errors.New(restResponse.String())
	}
	return respData, err
}
