package dao

import (
	"go-oauth2-server/common"
	"go-oauth2-server/db"
	"go-oauth2-server/model"
	"go-oauth2-server/util/timeUtil"
)

type IoauthCodeDao interface {
	Save(code *model.OauthCode) error
	QueryNotExpiredByUserId(userId int64) *model.OauthCode
	QueryNotExpiredByCode(code string) *model.OauthCode
}

type OauthCodeDaoImpl struct{}

func (codeDao *OauthCodeDaoImpl) Save(code *model.OauthCode) error {
	return db.GetDb().Save(code).Error
}

func (codeDao *OauthCodeDaoImpl) QueryNotExpiredByCode(code string) *model.OauthCode {
	oauthCode := &model.OauthCode{}
	err := db.GetDb().Find(&oauthCode, "code = ? and expired_at > ? and del = 0", code, timeUtil.GetNowTimestamp()).Error
	if err != nil {
		if err.Error() == common.DB_RECORD_NOT_EXIST {
			return nil
		} else {
			panic(common.Failure(common.DB_QUERY_ERROR))
		}
	}
	return oauthCode
}

func (codeDao *OauthCodeDaoImpl) QueryNotExpiredByUserId(userId int64) *model.OauthCode {
	oauthCode := &model.OauthCode{}
	err := db.GetDb().Find(&oauthCode, "user_id = ? and expired_at > ? and del = 0", userId, timeUtil.GetNowTimestamp()).Error
	if err != nil {
		if err.Error() == common.DB_RECORD_NOT_EXIST {
			return nil
		} else {
			panic(common.Failure(common.DB_QUERY_ERROR))
		}
	}
	return oauthCode
}
