package timeUtil

import "time"

//获取当前时间戳(毫秒)
func GetNowTimestamp() int64 {
	return time.Now().UnixNano() / 1e6
}
