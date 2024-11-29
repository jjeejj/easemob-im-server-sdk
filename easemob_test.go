package easemobimserversdk

import (
	"context"
	"easemob-im-server-sdk/config"
	"easemob-im-server-sdk/token"
	"encoding/json"
	"testing"
)

var easemobClient *Easemob

func init() {
	easemobClient = New(&config.EasemobConfig{
		ApiUrl:       "http://a1.easecdn.com",
		OrgName:      "xxxx",
		AppName:      "demo",
		AppKey:       "xxxx#demo",
		ClientId:     "xxxx",
		ClientSecret: "xxxx",
	})
}

var appToken = "YWMtuGQtRK4sEe-E6uULlrs2wpgXHToj1TUqkTguT5AFcZal563Tg2FHiqkrRPj0_wlqAgMAAAGTdw52CgAADhAGhI5ZZx6ehVmutB7GP4xFO3hspcXtk8IxJ_CgOtGryQ"

func TestGetAppToken(t *testing.T) {
	tokenResp, err := easemobClient.Token.GetAppToken(context.Background(), &token.GetAppTokenReq{
		TTl: 3600,
	})
	if err != nil {
		t.Error(err)
	}
	tokenRespByte, _ := json.Marshal(tokenResp)
	t.Logf("tokenResp %v", string(tokenRespByte))
}

func TestGetUserToken(t *testing.T) {
	tokenResp, err := easemobClient.Token.GetUserToken(context.Background(), &token.GetUserTokenReq{
		TTl:            3600,
		Username:       "test_21212121212",
		AutoCreateUser: true,
		AppToken:       appToken,
	})
	if err != nil {
		t.Error(err)
	}
	tokenRespByte, _ := json.Marshal(tokenResp)
	t.Logf("tokenResp %v", string(tokenRespByte))
}

func TestGetUserTokenByPassword(t *testing.T) {
	tokenResp, err := easemobClient.Token.GetUserTokenByPassword(context.Background(), &token.GetUserTokenByPasswordReq{
		TTl:      3600,
		Username: "87352467",
		Password: "123456",
	})
	if err != nil {
		t.Error(err)
	}
	tokenRespByte, _ := json.Marshal(tokenResp)
	t.Logf("tokenResp %v", string(tokenRespByte))
}

func TestGetUserDynamicToken(t *testing.T) {
	tokenResp, err := easemobClient.Token.GetUserDynamicToken(context.Background(), &token.GetUserDynamicTokenReq{
		TTl:      3600,
		Username: "87352467",
	})
	if err != nil {
		t.Error(err)
	}
	tokenRespByte, _ := json.Marshal(tokenResp)
	t.Logf("tokenResp %v", string(tokenRespByte))
}
