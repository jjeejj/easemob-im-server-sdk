package easemobimserversdk

import (
	"context"
	"easemob-im-server-sdk/token"
	"testing"
)

var easemobClient *Easemob

func init() {
	easemobClient = New(&EasemobConfig{
		OrgName:      "easemob-cn",
		AppName:      "test",
		AppKey:       "easemob-cn-test",
		ClientId:     "easemob-cn-test",
		ClientSecret: "easemob-cn-test",
	})
}

func TestGetAppToken(t *testing.T) {
	easemobClient.Token.GetAppToken(context.Background(), &token.GetAppTokenReq{
		TTl: 3600,
	})
}
