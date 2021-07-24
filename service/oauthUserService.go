package service

import (
	"errors"
	"go-oauth2-server/common"
	"go-oauth2-server/dao"
	"go-oauth2-server/model"
	"go-oauth2-server/util/mymd5"
)

type IoauthUserService interface {
	Login(user *model.OauthUser) (*model.OauthUser, error)
}

type OauthUserServiceImpl struct {
	userDao dao.IoauthUserDao
}

func NewOauthUserServiceImpl(userDao dao.IoauthUserDao) IoauthUserService {
	return &OauthUserServiceImpl{
		userDao: userDao,
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
