package dao

import "go-oauth2-server/model"

type Ioauth2ClientDetailDao interface {
	Save(oauthClientDetail model.OauthClientDetail) error
}
