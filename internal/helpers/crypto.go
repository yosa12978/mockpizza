package helpers

import (
	"crypto/md5"
	"fmt"
)

func GetMD5Hash(s string) string {
	hash := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", hash)
}
