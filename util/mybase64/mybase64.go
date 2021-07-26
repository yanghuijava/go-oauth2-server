package mybase64

import (
	"encoding/base64"
	"errors"
	"strings"
)

func Decode(data string) (result string, err error) {
	if data == "" {
		return result, errors.New("字符串为空")
	}
	data = strings.Replace(data, " ", "", -1)
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return result, err
	}
	decodestr := string(decoded)
	return decodestr, nil
}
