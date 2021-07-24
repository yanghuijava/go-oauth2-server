package dto

type OauthAuthorizeRequest struct {
	ClientId     string
	RedirectUri  string
	ResponseType string
	Scope        string
	State        string
}
