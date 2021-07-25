package model

import "time"

type OauthAccessToken struct {
	Id         int64     `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Token      string    `gorm:"column:token;NOT NULL"`                                 // token
	ClientId   string    `gorm:"column:client_id;NOT NULL"`                             // 客户端Id
	UserId     int64     `gorm:"column:user_id;NOT NULL"`                               // 用户ID
	ExpiredAt  int64     `gorm:"column:expired_at;NOT NULL"`                            // 过期时间戳
	Del        int       `gorm:"column:del;default:0;NOT NULL"`                         //是否删除  0：否  1：是
	CreateTime time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *OauthAccessToken) TableName() string {
	return "oauth_access_token"
}
