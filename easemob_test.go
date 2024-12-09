package easemobimserversdk

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/jjeejj/easemob-im-server-sdk/chatroom"
	"github.com/jjeejj/easemob-im-server-sdk/config"
	"github.com/jjeejj/easemob-im-server-sdk/token"
	"github.com/jjeejj/easemob-im-server-sdk/user"
)

var easemobClient *Easemob

func init() {
	easemobClient, _ = New(&config.EasemobConfig{
		ApiUrl:       "http://a1.easecdn.com",
		OrgName:      "xxxx",
		AppName:      "demo",
		AppKey:       "xxxx#demo",
		ClientId:     "xxxx",
		ClientSecret: "xxxx",
	})
}

var appToken = "YWMt3Vtn_rXhEe-sJW-YODrWhZgXHToj1TUqkTguT5AFcZal563Tg2FHiqkrRPj0_wlqAgMAAAGTqZGpAwAADhDYHsWXJ6UOswc_juwMG0GiI4k9mR_7nihVzaSpg"

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

func TestCreateChatroom(t *testing.T) {
	tokenResp, err := easemobClient.ChatRoom.Create(context.Background(), &chatroom.CreateReq{
		Name:        "测试聊天室",
		Description: "测试聊天室描述",
		Owner:       "87352467",
		AppToken:    appToken,
		MaxUsers:    1000,
	})
	if err != nil {
		t.Error(err)
	}
	tokenRespByte, _ := json.Marshal(tokenResp)
	t.Logf("tokenResp %v", string(tokenRespByte))
}

func TestRegistryUser(t *testing.T) {
	registryUserResp, err := easemobClient.User.Registry(context.Background(), &user.RegistryReq{
		User: user.RegistryReqUserInfo{
			Username: "64899890",
			Password: "64899890",
		},
		AppToken: appToken,
	})
	if err != nil {
		t.Error(err)
	}
	registryUserRespByte, _ := json.Marshal(registryUserResp)
	t.Logf("registryUserResp %v", string(registryUserRespByte))
}

func TestBatchRegistryUser(t *testing.T) {
	registryUserResp, err := easemobClient.User.BatchRegistry(context.Background(), &user.BatchRegistryReq{
		Users: []user.RegistryReqUserInfo{
			{Username: "22474669", Password: "22474669"},
			{Username: "89425812", Password: "89425812"},
		},
		AppToken: appToken})
	if err != nil {
		t.Error(err)
	}
	registryUserRespByte, _ := json.Marshal(registryUserResp)
	t.Logf("registryUserResp %v", string(registryUserRespByte))
}
