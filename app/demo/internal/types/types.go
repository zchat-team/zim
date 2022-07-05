package types

type Msg struct {
	ConvType      int    `json:"conv_type,omitempty"`
	MsgType       int    `json:"msg_type,omitempty"`
	Sender        string `json:"sender,omitempty"`
	Target        string `json:"target,omitempty"`
	Content       string `json:"content,omitempty"`
	IsTransparent bool   `json:"is_transparent,omitempty"`
}

type ContactNotify struct {
}

// GroupNotify 群通知
type GroupNotify struct {
	Operator  string `json:"operator"`
	Operation string `json:"operation"` // "create": 创建群 "add": 添加群成员
	Data      string `json:"data"`
}

type AddItem struct {
	OperatorName string    `json:"operator_name"`
	TargetList   []*Target `json:"target_list"`
}

type Target struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type CreateReq struct {
	Owner   string   `json:"owner,omitempty"`
	Members []string `json:"members,omitempty"`
	Name    string   ` json:"name,omitempty"`
	GroupId string   `json:"group_id,omitempty"`
	Notice  string   `json:"notice,omitempty"`
	Intro   string   `json:"intro,omitempty"`
	Avatar  string   `json:"avatar,omitempty"`
}

type CreateRsp struct {
	GroupId string `json:"group_id,omitempty"`
}

type JoinReq struct {
	Uin     string `json:"uin,omitempty"`
	GroupId string `json:"group_id,omitempty"`
}

type AddReq struct {
	GroupId string   `json:"group_id,omitempty"`
	Members []string `json:"members,omitempty"`
}

type AddRsp struct {
}
