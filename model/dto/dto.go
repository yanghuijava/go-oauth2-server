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
	ClientId     string           `json:"clientId"`
	Secret       string           `json:"secret"`
	Code         string           `json:"code"`
	GrantType    string           `json:"grantType"`
	OauthUser    *model.OauthUser `json:"oauthUser"`
	Scope        string           `json:"scope"`
	RefreshToken string           `json:"refreshToken"`
	BasicAuth    string           `json:"basicAuth"`

	UserName string `json:"userName"`
	Password string `json:"password"`
}

type AccessTokenRespose struct {
	AccessToken  string `json:"accessToken,omitempty"`
	ExpiresIn    int    `json:"expiresIn,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	Openid       string `json:"openid,omitempty"`
	Scope        string `json:"scope,omitempty"`
}

type UserInfoResponse struct {
	NickName string `json:"nickName"`
	Nation   string `json:"nation"`
	Province string `json:"province"`
	City     string `json:"city"`
	Openid   string `json:"openid"`
}
