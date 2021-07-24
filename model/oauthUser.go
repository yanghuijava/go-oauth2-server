package model

import "time"

type OauthUser struct {
	Id         int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Name       string    `gorm:"column:name;NOT NULL"`                                  // 用户名
	Password   string    `gorm:"column:password;NOT NULL"`                              // 密码，md5加密
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *OauthUser) TableName() string {
	return "oauth_user"
}
