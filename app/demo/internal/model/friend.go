package model

import (
	"time"
)

type Friend struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement;comment:系统编号"`
	Uid       int64     `json:"uid" gorm:"size:64;not null;comment:用户ID"`
	FriendUid int64     `json:"friend_uid" gorm:"size:64;not null;comment:好友ID"`
	CreatedAt time.Time `json:"created_at" gorm:"size:3"`
	UpdatedAt time.Time `json:"updated_at" gorm:"size:3"`
	Alias     string    `json:"alias" gorm:"size:64;not null;default:'';comment:别名"`
}

func (_ *Friend) TableName() string {
	return "friend"
}
