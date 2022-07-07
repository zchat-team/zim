package model

import (
	"time"
)

// 登录日志
type UserLoginLog struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement"`      // 系统编号
	Uid       int64     `json:"uid" gorm:"size:64;not null;default:0"`   // 用户ID
	Ua        string    `json:"ua" gorm:"size:64;not null"`              // UserAgent
	DeviceId  string    `json:"device_id" gorm:"size:64;not null"`       // 设备ID
	Type      int       `json:"type" gorm:"size:8;not null;default:0"`   // 类型[1:登录;2:登出]
	Status    int       `json:"status" gorm:"size:8;not null;default:0"` // 状态[1:成功;2:失败]
	LoginIp   string    `json:"login_ip" gorm:"size:64;not null"`        // 登录IP
	CreatedAt time.Time `json:"created_at" gorm:"comment:创建时间"`
}

func (_ *UserLoginLog) TableName() string {
	return "user_login_log"
}
