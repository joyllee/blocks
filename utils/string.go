package utils

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"runtime"
)

// 生成32位MD5
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}

//用于打印panic时的堆栈
func GetStack() []byte {
	buf := make([]byte, 1<<12) //16kb
	num := runtime.Stack(buf, false)

	return buf[:num]
}

//IsExist ...
func IsExist(filepath string) bool {
	_, err := os.Stat(filepath)
	if err == nil {
		return true
	}
	return !os.IsNotExist(err)
}

//IsDir ...
func IsDir(filepath string) bool {
	f, err := os.Stat(filepath)
	if err != nil {
		return false
	}
	return f.IsDir()
}