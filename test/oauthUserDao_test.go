package test

import (
	"go-oauth2-server/dao"
	"go-oauth2-server/db"
	"go-oauth2-server/model"
	"go-oauth2-server/util/mymd5"
	"testing"
)

func TestUserSave(t *testing.T) {
	db.InitDb()
	defer db.Close()
	user := &model.OauthUser{Name: "admin", Password: mymd5.Md5("888888")}
	userDao := &dao.OauthUserDaoImpl{}
	err := userDao.Save(user)
	if err != nil {
		t.Error(err)
	}
}
