package dao

import (
	"go-oauth2-server/common"
	"go-oauth2-server/db"
	"go-oauth2-server/model"
)

type IoauthClientDetailDao interface {
	Save(oauthClientDetail *model.OauthClientDetail) error
	QueryByClientId(clientId string) *model.OauthClientDetail
}

type OauthClientDetailDaoImpl struct{}

func (oauthClientDetailDao *OauthClientDetailDaoImpl) Save(oauthClientDetail *model.OauthClientDetail) error {
	return db.GetDb().Save(oauthClientDetail).Error
}

func (oauth2ClientDetailDao *OauthClientDetailDaoImpl) QueryByClientId(clientId string) *model.OauthClientDetail {
	var detail = new(model.OauthClientDetail)
	err := db.GetDb().Find(&detail, "client_id=?", clientId).Error
	if err != nil {
		if err.Error() == common.DB_RECORD_NOT_EXIST {
			return nil
		} else {
			panic(common.Failure(common.DB_QUERY_ERROR))
		}
	}
	return detail
}
