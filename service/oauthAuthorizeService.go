package service

import (
	"errors"
	"github.com/sirupsen/logrus"
	"go-oauth2-server/common"
	"go-oauth2-server/common/err"
	"go-oauth2-server/common/grantType"
	"go-oauth2-server/dao"
	"go-oauth2-server/model"
	"go-oauth2-server/model/dto"
	"go-oauth2-server/util/mybase64"
	"go-oauth2-server/util/mymd5"
	"go-oauth2-server/util/myuuid"
	"go-oauth2-server/util/timeUtil"
	"strconv"
	"strings"
)

type IoauthAuthorizeService interface {
	//校验客户端参数
	CheckAuthorizeParams(request *dto.OauthAuthorizeRequest) error
	//获取授权码
	AuthorizeCode(request *dto.OauthAuthorizeRequest, user *model.OauthUser) (string, error)
	//授权码简化模式，直接返回token，而不是code,不支持refreshToken
	AuthorizeToken(request *dto.OauthAuthorizeRequest, user *model.OauthUser) (*dto.AccessTokenRespose, error)
	//获取accessToken
	AccessToken(request *dto.AccessTokenReuqest) (*dto.AccessTokenRespose, err.Err)
}

type OauthAuthorizeServiceImpl struct {
	clientDetailDao       dao.IoauthClientDetailDao
	codeDao               dao.IoauthCodeDao
	accessRefreshTokenDao dao.IAccessRefreshTokenDao
	oauthUserDao          dao.IoauthUserDao
}

func NewOauthAuthorizeServiceImpl(clientDetailDao dao.IoauthClientDetailDao,
	codeDao dao.IoauthCodeDao,
	accessRefreshTokenDao dao.IAccessRefreshTokenDao,
	userDao dao.IoauthUserDao) IoauthAuthorizeService {
	return &OauthAuthorizeServiceImpl{
		clientDetailDao:       clientDetailDao,
		codeDao:               codeDao,
		accessRefreshTokenDao: accessRefreshTokenDao,
		oauthUserDao:          userDao,
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
	if !(request.ResponseType == common.RESP_TYPE_CODE || request.ResponseType == common.RESP_TYPE_TOKEN) {
		return errors.New("不支持的response_type值")
	}
	if request.RedirectUri != clientDetail.WebServerRedirectUri {
		return errors.New("redirect_uri参数错误")
	}
	return nil
}

//获取授权码
func (authorizeService *OauthAuthorizeServiceImpl) AuthorizeCode(request *dto.OauthAuthorizeRequest,
	user *model.OauthUser) (code string, err error) {
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

func (authorizeService *OauthAuthorizeServiceImpl) AuthorizeToken(request *dto.OauthAuthorizeRequest,
	user *model.OauthUser) (respose *dto.AccessTokenRespose, e error) {
	e = authorizeService.CheckAuthorizeParams(request)
	if e != nil {
		return nil, e
	}
	accessTokenReuqest := &dto.AccessTokenReuqest{
		ClientId:  request.ClientId,
		OauthUser: user,
		Scope:     request.Scope,
		GrantType: grantType.IMPLICIT,
	}
	respose, er := authorizeService.AccessToken(accessTokenReuqest)
	if er != nil {
		return nil, errors.New(er.Err().GetDesc())
	}
	return respose, nil
}

func (authorizeService *OauthAuthorizeServiceImpl) checkClient(request *dto.AccessTokenReuqest) (*model.OauthClientDetail, err.Err) {
	if request.ClientId == "" {
		return nil, err.NewErr(common.CLIENT_ID_EMPTY)
	}
	client := authorizeService.clientDetailDao.QueryByClientId(request.ClientId)
	if client == nil {
		return nil, err.NewErr(common.CLIENT_ID_NOT_EXIST)
	}
	if !client.IsExist(request.GrantType) {
		return nil, err.NewErr(common.CLIENT_NOT_SUPPORT)
	}
	return client, nil
}

func (authorizeService *OauthAuthorizeServiceImpl) AccessToken(request *dto.AccessTokenReuqest) (reponse *dto.AccessTokenRespose, e err.Err) {
	switch request.GrantType {
	case grantType.CODE:
		return handleCodeModel(request, authorizeService)
	case grantType.IMPLICIT:
		return handleImclicitModel(request, authorizeService)
	case grantType.REFRESH:
		return handleRefreshModel(request, authorizeService)
	case grantType.PASSWORD:
		return handlePassword(request, authorizeService)
	case grantType.CLIENT:
	default:
		e = err.NewErr(common.NO_NSUPPORT_GRANTTYPE)
		break
	}
	return reponse, e
}

func handlePassword(request *dto.AccessTokenReuqest,
	authorizeService *OauthAuthorizeServiceImpl) (reponse *dto.AccessTokenRespose, e err.Err) {
	if request.BasicAuth == "" {
		return nil, err.NewErr(common.AUTHORIZATION_EMPTY)
	}
	basicAuth := strings.Replace(request.BasicAuth, "Basic ", "", -1)
	data, er := mybase64.Decode(basicAuth)
	if er != nil {
		logrus.Errorf("base64解码错误：%s", er.Error())
		return nil, err.NewErr(common.BASE64_ERROR)
	}
	logrus.Infof("解码后数据：%s", data)
	dataArr := strings.Split(data, ":")
	if len(dataArr) != 2 {
		return nil, err.NewErr(common.PARAMS_ERROR)
	}
	request.ClientId = dataArr[0]
	request.Secret = dataArr[1]
	client, e := authorizeService.checkClient(request)
	if e != nil {
		return nil, e
	}
	if client.ClientSecret != request.Secret {
		return nil, err.NewErr(common.CLIENT_SECRET_ERROR)
	}
	userFind := authorizeService.oauthUserDao.QueryByName(request.UserName)
	if userFind == nil {
		return nil, err.NewErr(common.USER_NOT_EXIST)
	}
	if userFind.Password != mymd5.Md5(request.Password) {
		return nil, err.NewErr(common.PASSWORD_ERROR)
	}
	//生成token和refreshToken,如果存在就直接删除，产生新的
	accessTokenSave := &model.OauthAccessToken{
		Token:     myuuid.SimpleUUID(),
		ClientId:  client.ClientId,
		UserId:    userFind.Id,
		Scope:     request.Scope,
		ExpiredAt: timeUtil.GetNowTimestamp() + int64(client.AccessTokenValidity*1000),
	}
	refreshTokenSave := &model.OauthRefreshToken{
		RefreshToken: myuuid.SimpleUUID(),
		ClientId:     client.ClientId,
		UserId:       userFind.Id,
		Scope:        request.Scope,
		ExpiredAt:    timeUtil.GetNowTimestamp() + int64(client.RefreshTokenValidity*1000),
	}
	e = authorizeService.accessRefreshTokenDao.SaveToken(accessTokenSave, refreshTokenSave, nil)
	if e != nil {
		return nil, e
	}
	reponse = &dto.AccessTokenRespose{
		AccessToken:  accessTokenSave.Token,
		ExpiresIn:    client.AccessTokenValidity,
		Scope:        request.Scope,
		RefreshToken: refreshTokenSave.RefreshToken,
		//openid=MD5(clientId + userId)
		Openid: mymd5.Md5(client.ClientId + strconv.FormatInt(userFind.Id, 10)),
	}
	return reponse, nil
}

func handleRefreshModel(request *dto.AccessTokenReuqest,
	authorizeService *OauthAuthorizeServiceImpl) (reponse *dto.AccessTokenRespose, e err.Err) {
	client, e := authorizeService.checkClient(request)
	if e != nil {
		return nil, e
	}
	if client.ClientSecret != request.Secret {
		return nil, err.NewErr(common.CLIENT_SECRET_ERROR)
	}
	if request.RefreshToken == "" {
		return nil, err.NewErr(common.REFRESH_TOKEN_EMPTY)
	}
	refresh := authorizeService.accessRefreshTokenDao.QueryRefreshTokenByRefreshToken(request.RefreshToken)
	if refresh == nil {
		return nil, err.NewErr(common.REFRESH_TOKEN_INVALID)
	}
	access := authorizeService.accessRefreshTokenDao.QueryAccessTokenByUserId(refresh.UserId, refresh.ClientId)
	reponse = &dto.AccessTokenRespose{
		ExpiresIn:    client.AccessTokenValidity,
		Scope:        refresh.Scope,
		RefreshToken: refresh.RefreshToken,
		//openid=MD5(clientId + userId)
		Openid: mymd5.Md5(client.ClientId + strconv.FormatInt(refresh.UserId, 10)),
	}
	// accessToken存在且未过期，只延长过期时间，不生成新的accessToken；不存在生成新的
	if access == nil {
		accessTokenSave := &model.OauthAccessToken{
			Token:     myuuid.SimpleUUID(),
			ClientId:  refresh.ClientId,
			UserId:    refresh.UserId,
			ExpiredAt: timeUtil.GetNowTimestamp() + int64(client.AccessTokenValidity*1000),
			Scope:     refresh.Scope,
		}
		e := authorizeService.accessRefreshTokenDao.SaveAccessToken(accessTokenSave)
		if e != nil {
			return nil, e
		}
		reponse.AccessToken = accessTokenSave.Token
	} else {
		e := authorizeService.accessRefreshTokenDao.UpdateAccessTokenExpiredAtByToken(access.Token, timeUtil.GetNowTimestamp()+int64(client.AccessTokenValidity*1000))
		if e != nil {
			return nil, e
		}
		reponse.AccessToken = access.Token
	}
	return reponse, nil
}

//处理简化模式逻辑
func handleImclicitModel(request *dto.AccessTokenReuqest,
	authorizeService *OauthAuthorizeServiceImpl) (reponse *dto.AccessTokenRespose, e err.Err) {
	client, e := authorizeService.checkClient(request)
	if e != nil {
		return nil, e
	}
	if request.OauthUser == nil {
		return nil, err.NewErr(common.USER_NOT_AUTH)
	}
	//生成token和refreshToken,如果存在就直接删除，产生新的
	accessTokenSave := &model.OauthAccessToken{
		Token:     myuuid.SimpleUUID(),
		ClientId:  request.ClientId,
		UserId:    request.OauthUser.Id,
		Scope:     request.Scope,
		ExpiredAt: timeUtil.GetNowTimestamp() + int64(client.AccessTokenValidity*1000),
	}
	err := authorizeService.accessRefreshTokenDao.SaveToken(accessTokenSave, nil, nil)
	if err != nil {
		return nil, err
	}
	reponse = &dto.AccessTokenRespose{
		AccessToken: accessTokenSave.Token,
		ExpiresIn:   client.AccessTokenValidity,
		Scope:       request.Scope,
		//openid=MD5(clientId + userId)
		Openid: mymd5.Md5(client.ClientId + strconv.FormatInt(request.OauthUser.Id, 10)),
	}
	return reponse, nil
}

//处理授权码模式逻辑
func handleCodeModel(request *dto.AccessTokenReuqest,
	authorizeService *OauthAuthorizeServiceImpl) (reponse *dto.AccessTokenRespose, e err.Err) {
	client, e := authorizeService.checkClient(request)
	if e != nil {
		return nil, e
	}
	if request.Code == "" {
		return nil, err.NewErr(common.CODE_EMPTY)
	}
	if client.ClientSecret != request.Secret {
		return nil, err.NewErr(common.CLIENT_SECRET_ERROR)
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
		Scope:     oauthCode.Scope,
		ExpiredAt: timeUtil.GetNowTimestamp() + int64(client.AccessTokenValidity*1000),
	}
	refreshTokenSave := &model.OauthRefreshToken{
		RefreshToken: myuuid.SimpleUUID(),
		ClientId:     oauthCode.ClientId,
		UserId:       oauthCode.UserId,
		Scope:        oauthCode.Scope,
		ExpiredAt:    timeUtil.GetNowTimestamp() + int64(client.RefreshTokenValidity*1000),
	}
	//code使用一次必须删除
	oauthCode.Del = common.DEL
	err := authorizeService.accessRefreshTokenDao.SaveToken(accessTokenSave, refreshTokenSave, oauthCode)
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
