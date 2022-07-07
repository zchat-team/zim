package model

import (
	"time"
)

type Group struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement:false;comment:系统编号"`
	Owner     int64     `json:"owner" gorm:"size:64;not null"`         // 群主
	Type      int       `json:"type" gorm:"size:8;not null;default:0"` // 群类型
	Name      string    `json:"name" gorm:"size:64;not null"`          // 群名称
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"comment:更新时间"`
	Notice    string    `json:"notice" gorm:"size:256;not null;default:''"`
	Intro     string    `json:"intro" gorm:"size:256;not null;default:''"`
	Avatar    string    `json:"avatar" gorm:"size:256;not null;default:''"`
}

func (_ *Group) TableName() string {
	return "group"
}
