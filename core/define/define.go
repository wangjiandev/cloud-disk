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
