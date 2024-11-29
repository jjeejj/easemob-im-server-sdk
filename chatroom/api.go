package chatroom

const (
	// POST https://{host}/{org_name}/{app_name}/chatrooms
	ApiCreateChatroomPath = "/chatrooms" // 创建聊天室
)

type CreateReq struct {
	Name        string   `json:"name"`        // 聊天室名称，最大长度为 128 个字符
	Description string   `json:"description"` // 聊天室描述，最大长度为 512 个字符。
	MaxUsers    int      `json:"max_users"`   // 聊天室最大成员数（包括聊天室所有者）。取值范围为 [1,10,000]，默认值为 1000。如需调整请联系商务。
	Owner       string   `json:"owner"`       // 聊天室所有者
	Members     []string `json:"members"`     // 聊天室普通成员和管理员的用户 ID 数组，不包含聊天室所有者的用户 ID。该数组可包含的元素数量不超过 maxusers 的值。若传该参数，确保至少设置一个数组元素。
	Custom      string   // 聊天室扩展信息，例如，可以给聊天室添加业务相关的标记，不能超过 8 KB
}

type CreateResp struct {
	Data struct {
		Id string `json:"id"`
	} `json:"data"`
}
