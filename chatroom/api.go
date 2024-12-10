package chatroom

const (
	// POST https://{host}/{org_name}/{app_name}/chatrooms
	ApiCreateChatroomPath = "/chatrooms" // 创建聊天室
	// POST https://{host}/{org_name}/{app_name}/messages/chatrooms
	ApiSendMessagePath = "/messages/chatrooms" // 发送消息
)

type CreateReq struct {
	Name        string   `json:"name"`        // 聊天室名称，最大长度为 128 个字符
	Description string   `json:"description"` // 聊天室描述，最大长度为 512 个字符。
	MaxUsers    int      `json:"max_users"`   // 聊天室最大成员数（包括聊天室所有者）。取值范围为 [1,10,000]，默认值为 1000。如需调整请联系商务。
	Owner       string   `json:"owner"`       // 聊天室所有者
	Members     []string `json:"members"`     // 聊天室普通成员和管理员的用户 ID 数组，不包含聊天室所有者的用户 ID。该数组可包含的元素数量不超过 maxusers 的值。若传该参数，确保至少设置一个数组元素。
	Custom      string   // 聊天室扩展信息，例如，可以给聊天室添加业务相关的标记，不能超过 8 KB
	AppToken    string   `json:"app_token"` // app token
}

type CreateResp struct {
	Data struct {
		Id string `json:"id"`
	} `json:"data"`
}

// MessageType 消息类型
type MessageType string

const (
	MessageTypeText     MessageType = "txt"    // 文本消息
	MessageTypeImage    MessageType = "img"    //图片消息
	MessageTypeFile     MessageType = "file"   // 文件消息
	MessageTypeAudio    MessageType = "audio"  // 语音消息
	MessageTypeVideo    MessageType = "video"  // 视频消息
	MessageTypeLocation MessageType = "loc"    // 位置消息
	MessageTypeCmd      MessageType = "cmd"    // 透传消息
	MessageTypeCustom   MessageType = "custom" // 自定义消息
)

type SendMessageReq struct {
	AppToken         string      `json:"app_token"`          // app token
	From             string      `json:"from"`               // 消息发送方的用户 ID。若不传入该字段，服务器默认设置为 admin。
	To               []string    `json:"to"`                 // 消息接收方聊天室 ID 数组。每次最多可向 10 个聊天室发送消息
	ChatroomMsgLevel string      `json:"chatroom_msg_level"` // 聊天室消息优先级：- high：高；- （默认）normal：普通；- low：低
	Type             MessageType `json:"type"`               // 消息类型
	RoamIgnoreUsers  string      `json:"roam_ignore_users"`  // 设置哪些用户拉漫游消息时拉不到该消息。每次最多可传入 20 个用户 ID
	Ext              any         `json:"ext"`                // 消息支持扩展字段，可添加自定义信息。不能对该参数传入 null。同时，推送通知也支持自定义扩展字段，详见 APNs 自定义显示 和 Android 推送字段说明。
	Body             any         `json:"body"`               // 消息内容 json
}

// SendTextMessageBody 发送的文本消息 body
type SendTxtMessageBody struct {
	Msg string `json:"msg"`
}

// SendImgMessageBody 发送的图片消息 body
type SendImgMessageBody struct {
	FileName string         `json:"filename"` // 图片名称。建议传入该参数，否则客户端收到图片消息时无法显示图片名称
	Url      string         `json:"url"`      // 片 URL 地址：https://{host}/{org_name}/{app_name}/chatfiles/{file_uuid}。其中 file_uuid 为文件 ID，成功上传图片文件后，从 文件上传 的响应 body 中获取
	Size     map[string]int `json:"size"`     // 图片尺寸，单位为像素，包含以下字段：- height：图片高度；- width：图片宽度。
	Secret   string         `json:"secret"`   // 图片的访问密钥，即成功上传图片后，从 文件上传 的响应 body 中获取的 share-secret。如果图片文件上传时设置了文件访问限制（restrict-access），则该字段为必填
}

// SendImgMessageBody 发送的语音消息 body
type SendAudioMessageBody struct {
	FileName string `json:"filename"` // 语音文件的名称。建议传入该参数，否则客户端收到语音消息时无法显示语音文件名称
	Url      string `json:"url"`      // 语音文件 URL 地址：https://{host}/{org_name}/{app_name}/chatfiles/{file_uuid}。file_uuid 为文件 ID，成功上传语音文件后，从 文件上传 的响应 body 中获取。
	Length   int64  `json:"length"`   //语音时长，单位为秒。
	Secret   string `json:"secret"`   // 语音文件访问密钥，即成功上传语音文件后，从 文件上传 的响应 body 中获取的 share-secret。 如果语音文件上传时设置了文件访问限制（restrict-access），则该字段为必填
}

// SendVideoMessageBody 发送的视频消息 body
type SendVideoMessageBody struct {
	FileName    string `json:"filename"`     //视频文件名称。建议传入该参数，否则客户端收到视频消息时无法显示视频文件名称。
	Thumb       string `json:"thumb"`        // 视频缩略图 URL 地址：https://{host}/{org_name}/{app_name}/chatfiles/{file_uuid}。file_uuid 为视频缩略图唯一标识，成功上传缩略图文件后，从 文件上传 的响应 body 中获取。
	Length      int64  `json:"length"`       //视频时长，单位为秒。
	Url         string `json:"url"`          // 视频文件 URL 地址：https://{host}/{org_name}/{app_name}/chatfiles/{file_uuid}。其中 file_uuid 为文件 ID，成功上传视频文件后，从 文件上传 的响应 body 中获取。
	Secret      string `json:"secret"`       //视频文件访问密钥，即成功上传视频文件后，从 文件上传 的响应 body 中获取的 share-secret。如果视频文件上传时设置了文件访问限制（restrict-access），则该字段为必填。
	FileLength  int64  `json:"file_length"`  // 视频文件大小，单位为字节。
	ThumbSecret string `json:"thumb_secret"` // 视频缩略图访问密钥，即成功上传视频文件后，从 文件上传 的响应 body 中获取的 share-secret。如果缩略图文件上传时设置了文件访问限制（restrict-access），则该字段为必填。
}

// SendFileMessageBody 发送的文件消息 body
type SendFileMessageBody struct {
	FileName string `json:"filename"` // 文件名称。建议传入该参数，否则客户端收到文件消息时无法显示文件名称。
	Secret   string `json:"secret"`   // 文件访问密钥，即成功上传文件后，从 文件上传 的响应 body 中获取的 share-secret。如果文件上传时设置了文件访问限制（restrict-access），则该字段为必填。
	Url      string `json:"url"`      // 文件 URL 地址：https://{host}/{org_name}/{app_name}/chatfiles/{file_uuid}。其中 file_uuid 为文件 ID，成功上传视频文件后，从 文件上传 的响应 body 中获取。
}

// SendLocationMessageBody 发送的位置消息 body
type SendLocationMessageBody struct {
	Lat  string `json:"lat"`  // 位置的纬度，单位为度。
	Lng  string `json:"lng"`  // 位置的经度，单位为度。
	Addr string `json:"addr"` // 位置的文字描述。
}

// SendCmdMessageBody 发送的命令消息 body
type SendCmdMessageBody struct {
	Action string `json:"action"` // 命令内容。
}

// SendCustomMessageBody 发送的自定义消息 body
type SendCustomMessageBody struct {
	CustomEvent string            `json:"customEvent"` // 用户自定义的事件类型。该参数的值必须满足正则表达式 [a-zA-Z0-9-_/\.]{1,32}，长度为 1-32 个字符。
	CustomExts  map[string]string `json:"customExts"`  // 用户自定义的事件属性，类型必须是 Map<String,String>，最多可以包含 16 个元素。customExts 是可选的，不需要可以不传。
}
type SendMessageResp struct {
	Data map[string]any `json:"data"` // 回数据详情。该字段的值为包含聊天室 ID 和 发送的消息的 ID 的键值对。
	// 例如 "185145305923585": "1029545553039460728"，表示在 ID 为 184524748161025 的聊天室中发送了消息 ID 为 1029545553039460728 的消息。
}
