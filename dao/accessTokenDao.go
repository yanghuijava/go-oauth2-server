package dao

import (
	"github.com/sirupsen/logrus"
	"go-oauth2-server/common"
	"go-oauth2-server/common/err"
	"go-oauth2-server/db"
	"go-oauth2-server/model"
	"go-oauth2-server/util/timeUtil"
)

type IAccessRefreshTokenDao interface {
	QueryAccessTokenByToken(token string) *model.OauthAccessToken

	QueryAccessRefreshTokenByUserId(userId int64, clientId string) (*model.OauthAccessToken, *model.OauthRefreshToken)
	// 保存token
	// 1、删除当前用户下所有未过期的accessToken，refreshToken
	// 2、保存新的accessToken
	// 3、refreshToken 不为nil则保存
	// 4、codeDel 不为nil则删除
	SaveToken(accessTokenSave *model.OauthAccessToken, refreshTokenSave *model.OauthRefreshToken, codeDel *model.OauthCode) err.Err
}

type AccessRefreshTokenDaoImpl struct{}

func (accessRefreshTokenDao *AccessRefreshTokenDaoImpl) QueryAccessTokenByToken(token string) *model.OauthAccessToken {
	var accessToken model.OauthAccessToken
	err := db.GetDb().Find(&accessToken, "token = ? and expired_at > ? and del = 0", token, timeUtil.GetNowTimestamp()).Error
	if err != nil {
		if err.Error() == common.DB_RECORD_NOT_EXIST {
			return nil
		} else {
			panic(common.Failure(common.DB_QUERY_ERROR))
		}
	}
	return &accessToken
}

func (accessRefreshTokenDao *AccessRefreshTokenDaoImpl) QueryAccessRefreshTokenByUserId(userId int64, clientId string) (accessToken *model.OauthAccessToken, refreshToken *model.OauthRefreshToken) {
	err := db.GetDb().Find(&accessToken, "user_id = ? and client_id = ? and expired_at > ? and del = 0",
		userId, clientId, timeUtil.GetNowTimestamp()).Error
	if err != nil {
		if err.Error() == common.DB_RECORD_NOT_EXIST {
			accessToken = nil
		} else {
			panic(common.Failure(common.DB_QUERY_ERROR))
		}
	}
	err = db.GetDb().Find(&refreshToken, "user_id = ? and client_id = ? and expired_at > ? and del = 0",
		userId, clientId, timeUtil.GetNowTimestamp()).Error
	if err != nil {
		if err.Error() == common.DB_RECORD_NOT_EXIST {
			refreshToken = nil
		} else {
			panic(common.Failure(common.DB_QUERY_ERROR))
		}
	}
	return accessToken, refreshToken
}

func (accessRefreshTokenDao *AccessRefreshTokenDaoImpl) SaveToken(accessTokenSave *model.OauthAccessToken,
	refreshTokenSave *model.OauthRefreshToken, codeDel *model.OauthCode) (errResult err.Err) {
	tx := db.GetDb().Begin()
	if tx.Error != nil {
		logrus.Errorf("数据库开启事务错误：%s", tx.Error.Error())
		return err.NewErr(common.DB_TX_ERROR)
	}
	defer func() {
		if r := recover(); r != nil {
			logrus.Errorf("数据库错误：%s", r)
			tx.Rollback()
			errResult = err.NewErr(common.DB_ERROR)
		}
	}()
	//查询未过期的accessToken,存在则删除
	var accessTokens []model.OauthAccessToken
	if e := tx.Find(&accessTokens, "client_id = ? and user_id = ? and expired_at > ? and del = 0",
		accessTokenSave.ClientId, accessTokenSave.UserId, timeUtil.GetNowTimestamp()).Error; e != nil {
		panic(e.Error())
	}
	if accessTokens != nil && len(accessTokens) > 0 {
		for _, v := range accessTokens {
			if e := tx.Model(&v).Update("del", common.DEL).Error; e != nil {
				panic(e.Error())
			}
		}
	}
	//查询未过期的refreshToken,存在则删除
	var refreshTokens []model.OauthRefreshToken
	if e := tx.Find(&refreshTokens, "client_id = ? and user_id = ? and expired_at > ? and del = 0",
		accessTokenSave.ClientId, accessTokenSave.UserId, timeUtil.GetNowTimestamp()).Error; e != nil {
		panic(e.Error())
	}
	if refreshTokens != nil && len(refreshTokens) > 0 {
		for _, v := range refreshTokens {
			if e := tx.Model(&v).Update("del", common.DEL).Error; e != nil {
				panic(e.Error())
			}
		}
	}
	if e := tx.Save(accessTokenSave).Error; e != nil {
		panic("保存accessToken错误：" + e.Error())
	}
	if refreshTokenSave != nil {
		if e := tx.Save(refreshTokenSave).Error; e != nil {
			panic("保存refreshToken错误：" + e.Error())
		}
	}
	if codeDel != nil {
		if e := tx.Model(&codeDel).Update("del", codeDel.Del).Error; e != nil {
			panic("删除code错误：" + e.Error())
		}
	}
	if e := tx.Commit().Error; e != nil {
		panic("提交事务错误：" + e.Error())
	}
	return nil
}
