# 环信 im 接口封装

## 使用方式

### 安装

`go get github.com/jjeejj/easemob-im-server-sdk`

### 调用

```golang

// 实例化
easemobClient = New(&config.EasemobConfig{
    ApiUrl:       "https://a1.easemob.com",
    OrgName:      "easemob-cn",
    AppName:      "test",
    AppKey:       "easemob-cn-test",
    ClientId:     "easemob-cn-test",
    ClientSecret: "easemob-cn-test",
})
// 获取应用 token
easemobClient.Token.GetAppToken(context.Background(), &token.GetAppTokenReq{
    TTl: 3600,
})
```

## 目前已经实现功能

### Token
    [x] 获取 APP Token
    [x] 获取 User Token

### 聊天室
    [x] 创建聊天室