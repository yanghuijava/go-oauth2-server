package test

import (
	"go-oauth2-server/db"
	"go-oauth2-server/model"
	"go-oauth2-server/util/myuuid"
	"go-oauth2-server/util/timeUtil"
	"testing"
)

func TestDb(t *testing.T) {
	db.InitDb()
	defer db.Close()
	var oauthCode model.OauthCode
	db.GetDb().Find(&oauthCode, "code = ?", "IX7yR5")
	oauthCode.Del = 1
	db.GetDb().Model(&oauthCode).Update("del", oauthCode.Del)
}

func TestDb1(t *testing.T) {
	db.InitDb()
	defer db.Close()
	tx := db.GetDb().Begin()
	accessToken := &model.OauthAccessToken{
		Token:     myuuid.SimpleUUID(),
		ClientId:  "fOTQB8es4GBQwcsy",
		UserId:    1,
		ExpiredAt: timeUtil.GetNowTimestamp() + int64(7200),
	}
	if err := tx.Save(accessToken).Error; err != nil {
		panic(err)
	}
	tx.Commit()
}

func TestDb2(t *testing.T) {
	db.InitDb()
	defer db.Close()
	var user model.OauthUser
	if err := db.GetDb().Find(&user, "id = ?", 1).Error; err != nil {
		t.Error(err)
	}
	db.GetDb().Save(&model.OauthUser{
		Id:       user.Id,
		NickName: "正是那朵玫瑰",
	})
}
