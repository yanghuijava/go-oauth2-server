package grantType

type GrantType string

const (
	CODE     = "authorization_code"
	PASSWORD = "password"
	IMPLICIT = "implicit"
	CLIENT   = "client_credentials"
	REFRESH  = "refresh_token"
)
