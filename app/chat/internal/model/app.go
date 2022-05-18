package model

type App struct {
	Id        int64  `json:"id" gorm:"primaryKey;column:id;type:bigint(20)"`
	Name      string `json:"name" gorm:"column:name;type:varchar(50);not null"`
	AppSecret string `json:"app_secret" gorm:"column:app_secret;type:varchar(64);not null"`
}

func (_ *App) TableName() string {
	return "app"
}
