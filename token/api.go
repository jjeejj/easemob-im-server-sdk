package token

// 接口路径
const (
	// POST https://{host}/{org_name}/{app_name}/token
	ApiAppTokenPath = "/token" // 接口请求的路径
	// POST https://{host}/{org_name}/{app_name}/token
	ApiUserTokenPath = "/token"
)

type GetAppTokenReq struct {
	// GrantType    string `json:"grant_type"` // 授权方式。该参数设置为固定字符串 client_credentials，即客户端凭证模式。
	// ClientId     string `json:"client_id"`  // App 的 client_id，用于生成 app token 调用 REST API。详见 环信即时通讯云控制台的应用详情页面。
	// ClientSecret string `json:"client_secret"`
	// 	token 有效期，单位为秒。
	// - 若传入该参数，token 有效期以传入的值为准。
	// - 若不传该参数，以 环信即时通讯云控制台的用户认证页面的 token 有效期的设置为准。
	// - 若设置为 0，则 token 永久有效。
	// 注意：VIP 5 集群该参数单位为毫秒。
	TTl int64 `json:"ttl"`
}

type GetAppTokenResp struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`  // Token 有效时间，单位为秒，在有效期内不需要重复获取
	Application string `json:"application"` // 当前 App 的 UUID 值
}

type GetUserTokenReq struct {
	Username       string `json:"username"`  // 用户 ID。
	AppToken       string `json:"app_token"` // app token
	TTl            int64  `json:"ttl"`
	AutoCreateUser bool   `json:"auto_create_user"` // 当用户不存在时，是否自动创建用户。自动创建用户时，需保证授权方式（grant_type）必须为 inherit，API 请求 header 中使用 App token 进行鉴权
}

type GetUserTokenByPasswordReq struct {
	Username string `json:"username"` // 用户 ID。
	Password string `json:"password"`
	TTl      int64  `json:"ttl"`
}

type GetUserTokenResp struct {
	AccessToken string `json:"access_token"`
	// 	token 有效期，单位为秒。在有效期内无需重复获取。
	// 注意：VIP 5 集群该参数单位为毫秒
	ExpiresIn int64    `json:"expires_in"`
	User      struct { // 用户信息
		Username string `json:"username"`
		UUID     string `json:"uuid"`
		Created  int64  `json:"created"`  // 册用户的 Unix 时间戳，单位为毫秒
		Modified int64  `json:"modified"` // 最近一次修改用户信息的 Unix 时间戳，单位为毫秒
		// 		用户是否为活跃状态：
		// - true：用户为活跃状态。
		// - false：用户为封禁状态。如要使用已被封禁的用户账户，你需要调用解禁用户的 API对账号解除封禁。
		Activated bool `json:"activated"`
	} `json:"user"`
}

type GetUserDynamicTokenReq struct {
	Username string `json:"username"` // 用户 ID
	TTl      int64  `json:"ttl"`
}

type GetUserDynamicTokenResp struct {
	AccessToken string `json:"access_token"`
	// 	token 有效期，单位为秒。在有效期内无需重复获取。
	ExpiresIn int64 `json:"expires_in"`
}
