package helper

import (
	"crypto/md5"
	"fmt"
	"wangjiandev/cloud-disk/core/define"

	"github.com/golang-jwt/jwt/v4"
)

// 生成MD5
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// 生成token
func GenerateToken(id uint64, identity string, name string) (string, error) {
	uc := define.UserClaim{}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, uc)
	signedString, err := token.SignedString([]byte(define.Jwtkey))
	if err != nil {
		return "", err
	}
	return signedString, nil
}
