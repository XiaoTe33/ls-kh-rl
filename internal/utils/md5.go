package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5Encoded(pre string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(pre)))
}
