package model

import (
	"gorm.io/plugin/soft_delete"
)

type Msg struct {
	Id         int64                 `json:"id" gorm:"primaryKey;column:id;type:bigint(20)"`
	ConvType   int                   `json:"conv_type" gorm:"column:conv_type;type:tinyint(4);not null"`
	Content    string                `json:"content" gorm:"column:content;type:varchar(5000);not null"`
	Type       int                   `json:"type" gorm:"column:type;type:int(11);not null;default:0"`
	DeletedAt  soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:bigint(20);not null;default:0"`
	Sender     string                `json:"sender" gorm:"column:sender;type:varchar(50);not null"`
	Target     string                `json:"target" gorm:"column:target;type:varchar(50);not null"`
	AtUserList string                `json:"at_user_list" gorm:"column:at_user_list;type:varchar(5000);not null"`
	ReadTime   int64                 `json:"read_time" gorm:"column:read_time;type:bigint(20);not null;default:0"`
	SendTime   int64                 `json:"send_time" gorm:"column:send_time;type:bigint(20);not null;default:0"`
	ClientUuid string                `json:"client_uuid" gorm:"column:client_uuid;type:varchar(50);not null"`
}

func (_ *Msg) TableName() string {
	return "msg"
}
