package model

import (
	"strings"
	"time"
)

type OauthClientDetail struct {
	Id           int64  `gorm:"column:id;primary_key;AUTO_INCREMENT"`
	ClientId     string `gorm:"column:client_id;NOT NULL"`     // 客户端ID
	ClientSecret string `gorm:"column:client_secret;NOT NULL"` // 客户端访问密匙

	Scope                 string    `gorm:"column:scope;NOT NULL"`                                  // 客户端申请的权限范围，可选值包括read,write,trust;若有多个权限范围用逗号(,)分隔
	AuthorizedGrantTypes  string    `gorm:"column:authorized_grant_types;NOT NULL"`                 // 客户端支持的授权许可类型(grant_type)，可选值包括authorization_code,password,refresh_token,implicit,client_credentials,若支持多个授权许可类型用逗号(,)分隔
	WebServerRedirectUri  string    `gorm:"column:web_server_redirect_uri;NOT NULL"`                // 客户端重定向URI，当grant_type为authorization_code或implicit时, 在Oauth的流程中会使用并检查与数据库内的redirect_uri是否一致
	AccessTokenValidity   int       `gorm:"column:access_token_validity;default:7200;NOT NULL"`     // 设定客户端的access_token的有效时间值(单位:秒)，若不设定值则使用默认的有效时间值：2 * 60 * 60， 2小时
	RefreshTokenValidity  int       `gorm:"column:refresh_token_validity;default:2592000;NOT NULL"` // 设定客户端的refresh_token的有效时间值(单位:秒)，若不设定值则使用默认的有效时间值：30 * 24 * 60 * 60 ，30天
	AdditionalInformation string    `gorm:"column:additional_information"`                          // 预留字段
	CreateTime            time.Time `gorm:"column:create_time;default:CURRENT_TIMESTAMP;NOT NULL"`  // 创建时间
	UpdateTime            time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP;NOT NULL"`  // 更新时间
}

func NewOauthClientDetail(clientId string, clientSecret string,
	scope string, authorizedGrantTypes string,
	webServerRedirectUri string) *OauthClientDetail {
	return &OauthClientDetail{
		ClientId:              clientId,
		ClientSecret:          clientSecret,
		Scope:                 scope,
		AuthorizedGrantTypes:  authorizedGrantTypes,
		WebServerRedirectUri:  webServerRedirectUri,
		AdditionalInformation: "{}",
	}
}

func (m *OauthClientDetail) TableName() string {
	return "oauth_client_detail"
}

func (m *OauthClientDetail) IsExist(grantType string) bool {
	return strings.Contains(m.AuthorizedGrantTypes, grantType)
}
