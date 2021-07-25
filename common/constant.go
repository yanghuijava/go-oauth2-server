package common

const DB_RECORD_NOT_EXIST = "record not found"

//删除标识
const DEL = 1

//未删除标识
const NOT_DEL = 0

const RESP_TYPE_CODE = "code"
const RESP_TYPE_TOKEN = "token"

type ResultCode int

const (
	SUCESS                ResultCode = 0
	ERROR                 ResultCode = -1
	PARAMS_BINGDING_ERROR ResultCode = 100001
	NOT_LOGIN             ResultCode = 100002
	DB_QUERY_ERROR        ResultCode = 100003
	USER_NOT_EXIST        ResultCode = 100004
	PASSWORD_ERROR        ResultCode = 100005
	PARAMS_ERROR          ResultCode = 100006
	NO_NSUPPORT_GRANTTYPE ResultCode = 100007
	CLIENT_ID_NOT_EXIST   ResultCode = 100008
	CLIENT_SECRET_ERROR   ResultCode = 100009
	CODE_EMPTY            ResultCode = 100010
	CODE_ERROR            ResultCode = 100011
	DB_TX_ERROR           ResultCode = 100012
	DB_ERROR              ResultCode = 100013
	CLIENT_NOT_SUPPORT    ResultCode = 100014
	USER_NOT_AUTH         ResultCode = 100015
	TOKEN_INVALID         ResultCode = 100016
	TOKEN_EMPTY           ResultCode = 100017
	REFRESH_TOKEN_EMPTY   ResultCode = 100018
	REFRESH_TOKEN_INVALID ResultCode = 100019
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
	case PARAMS_ERROR:
		return "参数错误"
	case NO_NSUPPORT_GRANTTYPE:
		return "不支持的授权模式"
	case CLIENT_ID_NOT_EXIST:
		return "client_id不存在"
	case CODE_EMPTY:
		return "授权码模式下，code不能为空"
	case CODE_ERROR:
		return "授权码不正确或者已失效"
	case CLIENT_SECRET_ERROR:
		return "secret不正确"
	case DB_TX_ERROR:
		return "数据库开启事务错误"
	case DB_ERROR:
		return "数据库错误"
	case CLIENT_NOT_SUPPORT:
		return "当前客户端不支持的授权模式"
	case USER_NOT_AUTH:
		return "用户未授权"
	case TOKEN_EMPTY:
		return "token不能为空"
	case TOKEN_INVALID:
		return "token无效"
	case REFRESH_TOKEN_EMPTY:
		return "refreshToken不能为空"
	case REFRESH_TOKEN_INVALID:
		return "refreshToken无效或者不支持refreshToken机制"
	default:
		return "未知"
	}
}
