package service

import (
	"errors"
	"go-oauth2-server/common"
	"go-oauth2-server/common/err"
	"go-oauth2-server/common/grantType"
	"go-oauth2-server/dao"
	"go-oauth2-server/model"
	"go-oauth2-server/model/dto"
	"go-oauth2-server/util/mymd5"
	"go-oauth2-server/util/myuuid"
	"go-oauth2-server/util/timeUtil"
	"strconv"
)

type IoauthAuthorizeService interface {
	//校验客户端参数
	CheckAuthorizeParams(request *dto.OauthAuthorizeRequest) error
	//获取授权码
	AuthorizeCode(request *dto.OauthAuthorizeRequest, user model.OauthUser) (string, error)

	AccessToken(request *dto.AccessTokenReuqest) (*dto.AccessTokenRespose, err.Err)
}

type OauthAuthorizeServiceImpl struct {
	clientDetailDao       dao.IoauthClientDetailDao
	codeDao               dao.IoauthCodeDao
	accessRefreshTokenDao dao.IAccessRefreshTokenDao
}

func NewOauthAuthorizeServiceImpl(clientDetailDao dao.IoauthClientDetailDao,
	codeDao dao.IoauthCodeDao,
	accessRefreshTokenDao dao.IAccessRefreshTokenDao) IoauthAuthorizeService {
	return &OauthAuthorizeServiceImpl{
		clientDetailDao:       clientDetailDao,
		codeDao:               codeDao,
		accessRefreshTokenDao: accessRefreshTokenDao,
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
	codeFind := authorizeService.codeDao.QueryNotExpiredByUserId(user.Id)
	if codeFind != nil {
		return codeFind.Code, nil
	}
	code = RandStringBytesMaskImprSrcUnsafe(6)
	now := timeUtil.GetNowTimestamp()
	expiredAt := now + 5*60*1000 //code 5分钟过期
	oauthCode := &model.OauthCode{
		Code:      code,
		UserId:    user.Id,
		ClientId:  request.ClientId,
		ExpiredAt: expiredAt,
		Scope:     request.Scope,
	}
	err = authorizeService.codeDao.Save(oauthCode)
	if err != nil {
		return code, err
	}
	return code, nil
}

func (authorizeService *OauthAuthorizeServiceImpl) AccessToken(request *dto.AccessTokenReuqest) (reponse *dto.AccessTokenRespose, e err.Err) {
	client := authorizeService.clientDetailDao.QueryByClientId(request.ClientId)
	if client == nil {
		return nil, err.NewErr(common.CLIENT_ID_NOT_EXIST)
	}
	if client.ClientSecret != request.Secret {
		return nil, err.NewErr(common.CLIENT_SECRET_ERROR)
	}
	switch request.GrantType {
	case grantType.CODE:
		return handleCodeModel(request, client, authorizeService)
	case grantType.IMPLICIT:
	case grantType.CLIENT:
	case grantType.PASSWORD:
	case grantType.REFRESH:
	default:
		e = err.NewErr(common.NO_NSUPPORT_GRANTTYPE)
		break
	}
	return reponse, e
}

//处理授权码模式逻辑
func handleCodeModel(request *dto.AccessTokenReuqest,
	client *model.OauthClientDetail,
	authorizeService *OauthAuthorizeServiceImpl) (reponse *dto.AccessTokenRespose, e err.Err) {
	if request.Code == "" {
		return nil, err.NewErr(common.CODE_EMPTY)
	}
	oauthCode := authorizeService.codeDao.QueryNotExpiredByCode(request.Code)
	if oauthCode == nil {
		return nil, err.NewErr(common.CODE_ERROR)
	}
	//生成token和refreshToken,如果存在就直接删除，产生新的
	accessTokenSave := &model.OauthAccessToken{
		Token:     myuuid.SimpleUUID(),
		ClientId:  oauthCode.ClientId,
		UserId:    oauthCode.UserId,
		ExpiredAt: timeUtil.GetNowTimestamp() + int64(client.AccessTokenValidity*1000),
	}
	refreshTokenSave := &model.OauthRefreshToken{
		RefreshToken: myuuid.SimpleUUID(),
		ClientId:     oauthCode.ClientId,
		UserId:       oauthCode.UserId,
		ExpiredAt:    timeUtil.GetNowTimestamp() + int64(client.RefreshTokenValidity*1000),
	}
	//code使用一次必须删除
	oauthCode.Del = common.DEL
	err := authorizeService.accessRefreshTokenDao.SaveCodeModelToken(accessTokenSave, refreshTokenSave, oauthCode)
	if err != nil {
		return nil, err
	}
	reponse = &dto.AccessTokenRespose{
		AccessToken:  accessTokenSave.Token,
		RefreshToken: refreshTokenSave.RefreshToken,
		ExpiresIn:    client.AccessTokenValidity,
		Scope:        oauthCode.Scope,
		//openid=MD5(clientId + userId)
		Openid: mymd5.Md5(oauthCode.ClientId + strconv.FormatInt(oauthCode.UserId, 10)),
	}
	return reponse, nil
}
