package common

const DB_RECORD_NOT_EXIST = "record not found"

type ResultCode int

const (
	SUCESS                ResultCode = 0
	ERROR                 ResultCode = -1
	PARAMS_BINGDING_ERROR ResultCode = 100001
	NOT_LOGIN             ResultCode = 100002
	DB_QUERY_ERROR        ResultCode = 100003
	USER_NOT_EXIST        ResultCode = 100004
	PASSWORD_ERROR        ResultCode = 100005
)

func (code ResultCode) GetCode() int {
	return int(code)
}

func (code ResultCode) GetDesc() string {
	switch code {
	case SUCESS:
		return "成功"
	case ERROR:
		return "失败"
	case PARAMS_BINGDING_ERROR:
		return "参数绑定错误"
	case NOT_LOGIN:
		return "未登录"
	case DB_QUERY_ERROR:
		return "数据库查询错误"
	case USER_NOT_EXIST:
		return "用户不存在"
	case PASSWORD_ERROR:
		return "密码错误"
	default:
		return "未知"
	}
}
