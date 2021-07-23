package test

import (
	"fmt"
	dao2 "go-oauth2-server/dao"
	"go-oauth2-server/model"
	"testing"
)

func TestSave(t *testing.T) {
	dao := &dao2.Oauth2ClientDetailDaoImpl{}
	err := dao.Save(model.NewOauthClientDetail("dddede", "scxxsdsdwdww", "all",
		"authorization_code", "http://baidu.com"))
	if err != nil {
		t.Error(err)
	}
}

func TestQueryByClientId(t *testing.T) {
	dao := &dao2.Oauth2ClientDetailDaoImpl{}
	fmt.Println(dao.QueryByClientId(""))
}
