package define

import "github.com/golang-jwt/jwt/v4"

type UserClaim struct {
	Id       uint64
	Identity string
	Name     string
	jwt.StandardClaims
}

var Jwtkey = "cloud-disk-key"
var MailKey = "XJEPJMFVQLIHTWLF"

// 腾讯云的OSS对象存储密钥
var SecretID = "AKIDtc5HEwNXQ4QVJsGNBj13oIWT9YhZBWbP"
var SecretKey = "XonWy93i5B0Zrsc4GBqDLDzUjBQP7Zlg"
