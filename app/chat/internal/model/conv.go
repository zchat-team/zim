package model

import (
	"time"
)

// 会话
type Conv struct {
	Id           int64     `json:"id" gorm:"primaryKey;column:id;type:bigint(20)"`                                 // 系统编号
	Type         int       `json:"type" gorm:"column:type;type:tinyint(4);not null;default:0"`                     // 会话类型[1:单聊;2:群聊]
	SubType      int       `json:"sub_type" gorm:"column:sub_type;type:tinyint(4);not null;default:0"`             // 子类型
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;type:datetime;not null"`                     // 创建时间
	UpdatedAt    time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime;not null"`                     // 更新时间
	Owner        string    `json:"owner" gorm:"column:owner;type:varchar(50);not null"`                            // 会话所有者
	Target       string    `json:"target" gorm:"column:target;type:varchar(50);not null"`                          // 对方ID或群ID
	PeerLastRead int64     `json:"peer_last_read" gorm:"column:peer_last_read;type:bigint(20);not null;default:0"` // 对方最后读时间
	PeerLastRecv int64     `json:"peer_last_recv" gorm:"column:peer_last_recv;type:bigint(20);not null;default:0"` // 对方最后接收消息时间
	IsTop        string    `json:"is_top" gorm:"column:is_top;type:char(1);not null;default:0"`                    // 是否置顶
	IsMute       string    `json:"is_mute" gorm:"column:is_mute;type:char(1);not null;default:0"`                  // 消息免打扰
	Remark       string    `json:"remark" gorm:"column:remark;type:varchar(50);not null"`                          // 备注
}

func (_ *Conv) TableName() string {
	return "conv"
}
