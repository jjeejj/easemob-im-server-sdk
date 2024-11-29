package chatroom

import "context"

// IChatroom 聊天室 提供的接口能力
// https://doc.easemob.com/document/server-side/chatroom_manage.html#创建聊天室
type IChatroom interface {
	Create(ctx context.Context, reqParam *CreateReq) (*CreateResp, error) // 创建聊天室
}

type Chatroom struct {
}

func New() IChatroom {
	return &Chatroom{}
}

func (c *Chatroom) Create(ctx context.Context, reqParam *CreateReq) (*CreateResp, error) {
	return nil, nil
}
