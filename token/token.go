package token

import (
	"context"
	"crypto/sha256"
	"easemob-im-server-sdk/config"
	"easemob-im-server-sdk/request"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// https://doc.easemob.com/document/server-side/easemob_app_token.html#http-请求
type IToken interface {
	GetAppToken(ctx context.Context, reqParam *GetAppTokenReq) (*GetAppTokenResp, error)
	GetUserToken(ctx context.Context, reqParam *GetUserTokenReq) (*GetUserTokenResp, error)                      // 通过 App token 获取用户级别的 token                             // 获取 app 级别的 token
	GetUserTokenByPassword(ctx context.Context, reqParam *GetUserTokenByPasswordReq) (*GetUserTokenResp, error)  // 通过用户名称 获取用户级别的 token
	GetUserDynamicToken(ctx context.Context, reqParam *GetUserDynamicTokenReq) (*GetUserDynamicTokenResp, error) // 基于 AppKey、AppSecret 和 userId 生成 token
}

type Token struct {
	config     *config.EasemobConfig
	httpClient *request.HttpClient // 发送请求的客户端
}

func New(config *config.EasemobConfig, httpClient *request.HttpClient) IToken {
	return &Token{
		httpClient: httpClient,
		config:     config,
	}
}

func (t *Token) GetAppToken(ctx context.Context, reqParam *GetAppTokenReq) (*GetAppTokenResp, error) {
	var respData *GetAppTokenResp = &GetAppTokenResp{}
	restResponse, err := t.httpClient.Post(ctx, ApiUserTokenPath, map[string]any{
		"grant_type":    "client_credentials",
		"client_id":     t.config.ClientId,
		"client_secret": t.config.ClientSecret,
		"ttl":           reqParam.TTl,
	}, nil, respData)
	if restResponse.StatusCode() != http.StatusOK {
		return nil, errors.New(restResponse.String())
	}
	return respData, err
}

func (t *Token) GetUserToken(ctx context.Context, reqParam *GetUserTokenReq) (*GetUserTokenResp, error) {
	var respData *GetUserTokenResp = &GetUserTokenResp{}
	restResponse, err := t.httpClient.Post(ctx, ApiAppTokenPath, map[string]any{
		"grant_type":     "inherit",
		"username":       reqParam.Username,
		"autoCreateUser": reqParam.AutoCreateUser,
		"ttl":            reqParam.TTl,
	}, map[string]string{
		"Authorization": "Bearer " + reqParam.AppToken,
	}, respData)
	if restResponse.StatusCode() != http.StatusOK {
		return nil, errors.New(restResponse.String())
	}
	return respData, err
}
func (t *Token) GetUserTokenByPassword(ctx context.Context, reqParam *GetUserTokenByPasswordReq) (*GetUserTokenResp, error) {
	var respData *GetUserTokenResp = &GetUserTokenResp{}
	restResponse, err := t.httpClient.Post(ctx, ApiAppTokenPath, map[string]any{
		"grant_type": "password",
		"username":   reqParam.Username,
		"password":   reqParam.Password,
		"ttl":        reqParam.TTl,
	}, nil, respData)
	if restResponse.StatusCode() != http.StatusOK {
		return nil, errors.New(restResponse.String())
	}
	return respData, err
}

func (t *Token) GetUserDynamicToken(ctx context.Context, reqParam *GetUserDynamicTokenReq) (*GetUserDynamicTokenResp, error) {
	// 获取当前时间戳，单位为秒
	curTime := time.Now().Unix()
	signatureStr := fmt.Sprintf("%s%s%s%d%d%s", t.config.ClientId, t.config.AppKey, reqParam.Username, curTime, reqParam.TTl, t.config.ClientSecret)
	signatureHashByte := sha256.Sum256([]byte(signatureStr))
	signatureHashStr := fmt.Sprintf("%x", signatureHashByte)
	tokenJsonStr := fmt.Sprintf(`{"signature": "%s", "appkey": "%s","userId": "%s", "curTime":%d, "ttl": %d}`, signatureHashStr, t.config.AppKey, reqParam.Username, curTime, reqParam.TTl)
	accessToken := base64.StdEncoding.EncodeToString([]byte(tokenJsonStr))
	return &GetUserDynamicTokenResp{
		AccessToken: accessToken,
		ExpiresIn:   reqParam.TTl,
	}, nil
}
