package model

import (
	"time"
)

type Application struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement;comment:系统编号"`
	Uid       int64     `json:"uid" gorm:"size:64;not null;default:0;comment:用户ID"`
	FriendUid int64     `json:"friend_uid" gorm:"size:64;not null;default:0;comment:好友ID"`
	Status    int       `json:"status" gorm:"size:8;not null;default:0;comment:状态"` // 状态[1:审核中 2:通过 3:拒绝]
	IsRead    string    `json:"is_read" gorm:"size:1;not null;default:'0';comment:是否已读"`
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:更新时间"`
	ExpiresAt int64     `json:"expires_at" gorm:"size:64;not null;default:0;comment:过期时间"`
}

func (_ *Application) TableName() string {
	return "application"
}
