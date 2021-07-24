package dao

import (
	"go-oauth2-server/common"
	"go-oauth2-server/db"
	"go-oauth2-server/model"
)

type OauthCodeDao interface {
	Save(code *model.OauthCode) error
	QueryByCode(code string) *model.OauthCode
}

type OauthCodeDaoImpl struct{}

func (codeDao *OauthCodeDaoImpl) Save(code *model.OauthCode) error {
	return db.GetDb().Save(code).Error
}

func (codeDao *OauthCodeDaoImpl) QueryByCode(code string) *model.OauthCode {
	oauthCode := &model.OauthCode{}
	err := db.GetDb().Find(&oauthCode, "code=?", code).Error
	if err != nil {
		if err.Error() == common.DB_RECORD_NOT_EXIST {
			return nil
		} else {
			panic(common.Failure(common.DB_QUERY_ERROR))
		}
	}
	return oauthCode
}
