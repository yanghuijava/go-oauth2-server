package dao

import (
	"go-oauth2-server/common"
	"go-oauth2-server/db"
	"go-oauth2-server/model"
)

type IoauthUserDao interface {
	Save(user *model.OauthUser) error
	QueryByName(name string) *model.OauthUser
}

type OauthUserDaoImpl struct{}

func (userDao *OauthUserDaoImpl) Save(user *model.OauthUser) error {
	return db.GetDb().Save(user).Error
}

func (userDao *OauthUserDaoImpl) QueryByName(name string) *model.OauthUser {
	var user = new(model.OauthUser)
	err := db.GetDb().Find(&user, "name=?", name).Error
	if err != nil {
		if err.Error() == common.DB_RECORD_NOT_EXIST {
			return nil
		} else {
			panic(common.Failure(common.DB_QUERY_ERROR))
		}
	}
	return user
}
