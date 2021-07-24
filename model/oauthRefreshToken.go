package model

import "time"

type OauthRefreshToken struct {
	Id           int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	RefreshToken string    `gorm:"column:refresh_token;NOT NULL"`                         // refresh_token
	ClientId     int64     `gorm:"column:client_id;NOT NULL"`                             // 客户端Id
	UserId       int64     `gorm:"column:user_id;NOT NULL"`                               // 用户ID
	ExpiredAt    int64     `gorm:"column:expired_at;NOT NULL"`                            // 过期时间戳
	CreateTime   time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime   time.Time `gorm:"column:update_time;NOT NULL"`                           // 更新时间
}

func (m *OauthRefreshToken) TableName() string {
	return "oauth_refresh_token"
}
