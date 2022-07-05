package model

import (
	"time"
)

// 用户
type User struct {
	Id             int64     `json:"id" gorm:"primaryKey;autoIncrement:false"` // 用户ID
	Zid            string    `json:"zid" gorm:"size:64;not null"`              // 知聊号
	Password       string    `json:"password" gorm:"size:64;not null"`         // 密码
	CreatedAt      time.Time `json:"created_at" gorm:"size:3"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"size:3"`
	Status         int       `json:"status" gorm:"size:8;not null;default:0"`                        // 状态
	Nickname       string    `json:"nickname" gorm:"size:64;not null"`                               // 昵称
	Mobile         string    `json:"mobile" gorm:"size:64;not null"`                                 // 手机号
	MobileVerified int       `json:"mobile_verified" gorm:"size:8;not null;default:0"`               // 手机是否被验证
	Email          string    `json:"email" gorm:"size:64;not null"`                                  // 邮箱
	EmailVerified  int       `json:"email_verified" gorm:"size:8;not null;default:0"`                // 邮箱是否被验证
	Avatar         string    `json:"avatar" gorm:"size:256;not null"`                                // 头像
	Gender         int       `json:"gender" gorm:"size:8;not null;default:0"`                        // 性别[1:男;2:女]
	RegisterIp     string    `json:"register_ip" gorm:"size:64;not null"`                            // 注册IP
	LastLoginTime  time.Time `json:"last_login_time" gorm:"size:3"`                                  // 最近登录时间
	LastLoginIp    string    `json:"last_login_ip" gorm:"size:64;not null"`                          // 最近登录IP
	LoginIpLimit   string    `json:"login_ip_limit" gorm:"column:login_ip_limit;type:text;not null"` // 登录限制,JSON格式
}

func (_ *User) TableName() string {
	return "user"
}
