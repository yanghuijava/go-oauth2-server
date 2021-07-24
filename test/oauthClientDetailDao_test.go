package test

import (
	"fmt"
	dao2 "go-oauth2-server/dao"
	"go-oauth2-server/db"
	"go-oauth2-server/model"
	"go-oauth2-server/service"
	"testing"
)

func TestSave(t *testing.T) {
	db.InitDb()
	defer db.Close()
	dao := &dao2.OauthClientDetailDaoImpl{}
	clientProduce := &service.DefaultClientProduce{}
	clientId, clientSecret := clientProduce.ClientIdSecret()
	err := dao.Save(model.NewOauthClientDetail(
		clientId,
		clientSecret,
		"all",
		"authorization_code",
		"https://www.baidu.com"))
	if err != nil {
		t.Error(err)
	}
}

func TestQueryByClientId(t *testing.T) {
	db.InitDb()
	defer db.Close()
	dao := &dao2.OauthClientDetailDaoImpl{}
	fmt.Println(dao.QueryByClientId(""))
}
