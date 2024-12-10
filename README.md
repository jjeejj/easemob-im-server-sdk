# 环信 im 接口封装

## 结构设计
按照功能模块化进行设计，后续可以按照这样的规则 增加对应的接口
```golang
├── LICENSE
├── README.md
├── chatroom # 聊天室模块
│   ├── api.go # 聊天室相关结构体
│   └── chatroom.go # 接口封装
├── config # 配置模块
│   └── config.go
├── easemob.go # 主文件
├── easemob_test.go # 测试文件
├── go.mod
├── go.sum
├── request # 对外发送请求
│   └── request.go
└── token  # 获取相关 token 模块
│   ├── api.go
│   └── token.go
└── user  # 用户 模块
    ├── api.go
    └── user.go
```

## 使用方式

### 安装

`go get github.com/jjeejj/easemob-im-server-sdk`

> 这里如果拉取不下拉，可以指定最新 commit 重新拉取 go get github.com/jjeejj/easemob-im-server-sdk@{commit-hash}

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
- [x] 获取 APP Token
- [x] 获取 User Token

### 聊天室
- [x] 创建聊天室
- [x] 发送聊天室消息

### 用户
- [x] 授权注册单个用户
- [x] 批量授权注册用户

