package util

import (
	"crypto/md5"
	"encoding/hex"
)

//该方法用于上传时文件名格式化
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
