package test

import (
	"fmt"
	"go-oauth2-server/util"
	"go-oauth2-server/util/mymd5"
	"go-oauth2-server/util/myuuid"
	timeUtil2 "go-oauth2-server/util/timeUtil"
	"testing"
)

func TestGetNowTimestamp(t *testing.T) {
	fmt.Println(timeUtil2.GetNowTimestamp())
}

func TestMd5(t *testing.T) {
	fmt.Println(mymd5.Md5("123456"))
}

func TestUUID(t *testing.T) {
	fmt.Println(myuuid.UUID())
	fmt.Println(myuuid.SimpleUUID())
}

func TestParseUrlQuery(t *testing.T) {
	m := util.ParseUrlQuery("client_id=test_client_1&redirect_uri=https://www.example.com&response_type=code&scope=all&state=dsdsd")
	fmt.Println(m)
}
