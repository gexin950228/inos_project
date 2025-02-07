package utils

import (
	"crypto/md5"
	"fmt"
)

func EncryptPassword(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}
