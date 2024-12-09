package user

// 接口路径
const (
	// POST https://{host}/{org_name}/{app_name}/users
	ApiRegistryUserPath = "/users" // 接口请求的路径
	// POST https://{host}/{org_name}/{app_name}/users
	ApiBatchRegistryUserPath = "/users"
)

type RegistryReq struct {
	AppToken string `json:"app_token"`
	UserName string `json:"username"`
	Password string `json:"password"`
}

type RegistryResp struct {
	Entities []struct {
		Username  string `json:"username"`
		UUID      string `json:"uuid"`
		Created   int64  `json:"created"`   // 册用户的 Unix 时间戳，单位为毫秒
		Modified  int64  `json:"modified"`  // 最近一次修改用户信息的 Unix 时间戳，单位为毫秒
		Activated bool   `json:"activated"` // 		用户是否为活跃状态：		// - true：用户为活跃状态。		// - false：用户为封禁状态。如要使用已被封禁
		Type      string `json:"type"`
	} `json:"entities"`
}

type BatchRegistryReq struct {
	AppToken string `json:"app_token"`
	Users    []struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"users"`
}

type BatchRegistryResp struct {
	Entities []struct {
		Username  string `json:"username"`
		UUID      string `json:"uuid"`
		Created   int64  `json:"created"`   // 册用户的 Unix 时间戳，单位为毫秒
		Modified  int64  `json:"modified"`  // 最近一次修改用户信息的 Unix 时间戳，单位为毫秒
		Activated bool   `json:"activated"` // 		用户是否为活跃状态：		// - true：用户为活跃状态。		// - false：用户为封禁状态。如要使用已被封禁
		Type      string `json:"type"`
	} `json:"entities"`
}
