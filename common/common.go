package common

type Result struct {
	Code int         `json:"code"`
	Desc string      `json:"desc"`
	Data interface{} `json:"data"`
}

func (r Result) IsSucess() bool {
	if r.Code == 0 {
		return true
	}
	return false
}

func Failure(code ResultCode) *Result {
	return &Result{Code: code.GetCode(), Desc: code.GetDesc()}
}

func Error() *Result {
	return &Result{Code: ERROR.GetCode(), Desc: ERROR.GetDesc()}
}

func Success() *Result {
	return &Result{Code: SUCESS.GetCode(), Desc: SUCESS.GetDesc()}
}
