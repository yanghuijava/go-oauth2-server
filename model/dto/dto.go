package dto

import "go-oauth2-server/model"

type OauthAuthorizeRequest struct {
	ClientId     string
	RedirectUri  string
	ResponseType string
	Scope        string
	State        string
}

type AccessTokenReuqest struct {
	ClientId  string           `json:"clientId"`
	Secret    string           `json:"secret"`
	Code      string           `json:"code"`
	GrantType string           `json:"grantType"`
	OauthUser *model.OauthUser `json:"oauthUser"`
	Scope     string           `json:"scope"`
}

type AccessTokenRespose struct {
	AccessToken  string `json:"accessToken"`
	ExpiresIn    int    `json:"expiresIn"`
	RefreshToken string `json:"refreshToken"`
	Openid       string `json:"openid"`
	Scope        string `json:"scope"`
}

type UserInfoResponse struct {
	NickName string `json:"nickName"`
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	Openid   string `json:"openid"`
}
