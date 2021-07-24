package mymd5

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(data string) string {
	w := md5.New()
	io.WriteString(w, data)
	//将str写入到w中
	result := fmt.Sprintf("%x", w.Sum(nil))
	return result
}
