package service

import (
	"errors"
	"go-oauth2-server/common"
	"go-oauth2-server/common/err"
	"go-oauth2-server/dao"
	"go-oauth2-server/model"
	"go-oauth2-server/model/dto"
	"go-oauth2-server/util/mymd5"
	"strconv"
)

type IoauthUserService interface {
	Login(user *model.OauthUser) (*model.OauthUser, error)
	UserInfo(token string) (*dto.UserInfoResponse, err.Err)
}

type OauthUserServiceImpl struct {
	userDao               dao.IoauthUserDao
	accessRefreshTokenDao dao.IAccessRefreshTokenDao
}

func NewOauthUserServiceImpl(userDao dao.IoauthUserDao, accessRefreshTokenDao dao.IAccessRefreshTokenDao) IoauthUserService {
	return &OauthUserServiceImpl{
		userDao:               userDao,
		accessRefreshTokenDao: accessRefreshTokenDao,
	}
}

func (userService OauthUserServiceImpl) Login(userQuery *model.OauthUser) (user *model.OauthUser, err error) {
	userFind := userService.userDao.QueryByName(userQuery.Name)
	if userFind == nil {
		return nil, errors.New(common.USER_NOT_EXIST.GetDesc())
	}
	password := mymd5.Md5(userQuery.Password)
	if password != userFind.Password {
		return nil, errors.New(common.PASSWORD_ERROR.GetDesc())
	}
	return userFind, nil
}

func (userService OauthUserServiceImpl) UserInfo(token string) (*dto.UserInfoResponse, err.Err) {
	accessToken := userService.accessRefreshTokenDao.QueryAccessTokenByToken(token)
	if accessToken == nil {
		return nil, err.NewErr(common.TOKEN_INVALID)
	}
	u := userService.userDao.QueryById(accessToken.UserId)
	if u == nil {
		return nil, err.NewErr(common.USER_NOT_EXIST)
	}
	resp := &dto.UserInfoResponse{
		NickName: u.NickName,
		Nation:   u.Nation,
		Province: u.Province,
		City:     u.City,
		Openid:   mymd5.Md5(accessToken.ClientId + strconv.FormatInt(u.Id, 10)),
	}
	return resp, nil
}
