package model

import (
	"time"
)

// 群成员
type GroupMember struct {
	Id        int64     `json:"id" gorm:"primaryKey;column:id;type:bigint(20) auto_increment"` // 系统编号
	GroupId   string    `json:"group_id" gorm:"column:group_id;type:varchar(50);not null"`
	Member    string    `json:"member" gorm:"column:member;type:varchar(50);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;type:datetime(3);not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;type:datetime(3);not null"`
	//DeletedAt soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:bigint(20);not null;default:0"`
}

func (_ *GroupMember) TableName() string {
	return "group_member"
}
