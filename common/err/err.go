package err

import "go-oauth2-server/common"

type Err interface {
	Err() common.ResultCode
}

type ResultCodeErr struct {
	code common.ResultCode
}

func (codeErr *ResultCodeErr) Err() common.ResultCode {
	return codeErr.code
}

func NewErr(code common.ResultCode) Err {
	return &ResultCodeErr{code: code}
}
