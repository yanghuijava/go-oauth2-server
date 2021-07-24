package service

import (
	"errors"
	"go-oauth2-server/dao"
	"go-oauth2-server/model"
	"go-oauth2-server/model/dto"
	"go-oauth2-server/util/timeUtil"
)

type IoauthAuthorizeService interface {
	//校验客户端参数
	CheckAuthorizeParams(request *dto.OauthAuthorizeRequest) error
	//获取授权码
	AuthorizeCode(request *dto.OauthAuthorizeRequest, user model.OauthUser) (string, error)
}

type OauthAuthorizeServiceImpl struct {
	clientDetailDao dao.IoauthClientDetailDao
	codeDao         dao.OauthCodeDao
}

func NewOauthAuthorizeServiceImpl(clientDetailDao dao.IoauthClientDetailDao,
	codeDao dao.OauthCodeDao) IoauthAuthorizeService {
	return &OauthAuthorizeServiceImpl{
		clientDetailDao: clientDetailDao,
		codeDao:         codeDao,
	}
}

func (authorizeService *OauthAuthorizeServiceImpl) CheckAuthorizeParams(request *dto.OauthAuthorizeRequest) error {
	clientDetail := authorizeService.clientDetailDao.QueryByClientId(request.ClientId)
	if clientDetail == nil {
		return errors.New("client_id参数错误")
	}
	if clientDetail.Scope != request.Scope {
		return errors.New("不支持的scope值")
	}
	if request.ResponseType != "code" {
		return errors.New("不支持的response_type值")
	}
	if request.RedirectUri != clientDetail.WebServerRedirectUri {
		return errors.New("redirect_uri参数错误")
	}
	return nil
}

func (authorizeService *OauthAuthorizeServiceImpl) AuthorizeCode(request *dto.OauthAuthorizeRequest,
	user model.OauthUser) (code string, err error) {
	err = authorizeService.CheckAuthorizeParams(request)
	if err != nil {
		return code, err
	}
	code = RandStringBytesMaskImprSrcUnsafe(6)
	now := timeUtil.GetNowTimestamp()
	expiredAt := now + 5*60*1000 //code 5分钟过期
	oauthCode := &model.OauthCode{
		Code:      code,
		UserId:    user.Id,
		ClientId:  request.ClientId,
		ExpiredAt: expiredAt,
		Del:       0,
	}
	authorizeService.codeDao.Save(oauthCode)
	return code, nil
}
