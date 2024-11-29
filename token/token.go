package token

import "context"

// https://doc.easemob.com/document/server-side/easemob_app_token.html#http-请求
type IToken interface {
	GetAppToken(ctx context.Context, reqParam *GetAppTokenReq) (*GetAppTokenResp, error)
	GetUserToken(ctx context.Context, reqParam *GetUserTokenReq) (*GetUserTokenResp, error)           // 通过 App token 获取用户级别的 token                             // 获取 app 级别的 token
	GetUserTokenByPassword(ctx context.Context, reqParam *GetUserTokenReq) (*GetUserTokenResp, error) // 通过用户名称 获取用户级别的 token
	GetUserDynamicToken(ctx context.Context, reqParam *GetUserTokenReq) (*GetUserTokenResp, error)    // 基于 AppKey、AppSecret 和 userId 生成 token

}

type Token struct{}

func New() IToken {
	return &Token{}
}

func (t *Token) GetAppToken(ctx context.Context, reqParam *GetAppTokenReq) (*GetAppTokenResp, error) {
	return nil, nil
}

func (t *Token) GetUserToken(ctx context.Context, reqParam *GetUserTokenReq) (*GetUserTokenResp, error) {
	return nil, nil
}
func (t *Token) GetUserTokenByPassword(ctx context.Context, reqParam *GetUserTokenReq) (*GetUserTokenResp, error) {
	return nil, nil
}

func (t *Token) GetUserDynamicToken(ctx context.Context, reqParam *GetUserTokenReq) (*GetUserTokenResp, error) {
	return nil, nil
}
