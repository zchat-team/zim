package model

import (
	"time"
)

type GroupMember struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement;comment:系统编号"`
	GroupId   string    `json:"group_id" gorm:"size:64;not null;default:'';comment:群ID"`
	Member    string    `json:"member" gorm:"size:64;not null;default:'';comment:群成员"`
	CreatedAt time.Time `json:"created_at" gorm:"size:3"`
	UpdatedAt time.Time `json:"updated_at" gorm:"size:3"`
}

func (_ *GroupMember) TableName() string {
	return "group_member"
}
