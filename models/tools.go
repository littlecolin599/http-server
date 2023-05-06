package models

import (
	"crypto/md5"
	"fmt"
	"io"

	_ "github.com/go-sql-driver/mysql"
)

//工具包
//加密  MD5
func Md5(str string) string {
	//STR
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}
