package model

import (
	"time"
)

// 群成员
type GroupMember struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement;comment:系统编号"`
	GroupId   int64     `json:"group_id" gorm:"size:64;not null"` // 群ID
	Member    int64     `json:"member" gorm:"size:64;not null"`   // 成员用户ID
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:更新时间"`
}

func (_ *GroupMember) TableName() string {
	return "group_member"
}
