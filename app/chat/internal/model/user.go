package model

// 用户
type User struct {
	Id       int64  `json:"id" gorm:"primaryKey;column:id;type:bigint(20)"`
	Uin      string `json:"uin" gorm:"column:uin;type:varchar(50);not null"`
	Nickname string `json:"nickname" gorm:"column:nickname;type:varchar(50);not null"`
}

func (_ *User) TableName() string {
	return "user"
}
