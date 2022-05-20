package model

import (
	"gorm.io/plugin/soft_delete"
	"time"
)

// 群
type Group struct {
	Id           int64                 `json:"id" gorm:"primaryKey;column:id;type:bigint(20) auto_increment"` // 系统编号
	Owner        string                `json:"owner" gorm:"column:owner;type:varchar(50);not null"`
	GroupId      string                `json:"group_id" gorm:"column:group_id;type:varchar(50);not null"`
	Type         int                   `json:"type" gorm:"column:type;type:tinyint(4);not null;default:0"`
	Name         string                `json:"name" gorm:"column:name;type:varchar(50);not null"`
	CreatedAt    time.Time             `json:"created_at" gorm:"column:created_at;type:datetime(3)"`
	UpdatedAt    time.Time             `json:"updated_at" gorm:"column:updated_at;type:datetime(3)"`
	DeletedAt    soft_delete.DeletedAt `json:"deleted_at" gorm:"column:deleted_at;type:bigint(20);not null;default:0"`
	Notice       string                `json:"notice" gorm:"column:notice;type:varchar(200);not null"`
	Introduction string                `json:"introduction" gorm:"column:introduction;type:varchar(200);not null"`
	Avatar       string                `json:"avatar" gorm:"column:avatar;type:varchar(200);not null"`
}

func (_ *Group) TableName() string {
	return "group"
}
