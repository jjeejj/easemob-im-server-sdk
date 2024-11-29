package config

type EasemobConfig struct {
	ApiUrl       string // 接口请求的地址, 这里注意后缀不可以带有 /
	OrgName      string // 环信企业名
	AppName      string // 环信应用名
	AppKey       string // 环信应用key
	ClientId     string
	ClientSecret string
}
