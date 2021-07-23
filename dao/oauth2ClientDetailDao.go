package dao

import (
	"go-oauth2-server/db"
	"go-oauth2-server/model"
)

type Ioauth2ClientDetailDao interface {
	Save(oauthClientDetail *model.OauthClientDetail) error
	QueryByClientId(clientId string) *model.OauthClientDetail
}

type Oauth2ClientDetailDaoImpl struct{}

func (oauth2ClientDetailDao *Oauth2ClientDetailDaoImpl) Save(oauthClientDetail *model.OauthClientDetail) error {
	return db.GetDb().Save(oauthClientDetail).Error
}

func (oauth2ClientDetailDao *Oauth2ClientDetailDaoImpl) QueryByClientId(clientId string) *model.OauthClientDetail {
	var detail = new(model.OauthClientDetail)
	err := db.GetDb().Find(&detail, "client_id=?", clientId).Error
	if err != nil {
		if err.Error() == "record not found" {
			return nil
		} else {
			panic("查询异常：" + err.Error())
		}
	}
	return detail
}
