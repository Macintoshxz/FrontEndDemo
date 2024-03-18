package define

import "github.com/dgrijalva/jwt-go"

// token常量
type UserClaim struct {
	Id   int
	Name string
	Nick string
	jwt.StandardClaims
}

// token时效
var TokenExpire = 3600
var RefreshTokenExpire = 7200

// token Key
var JwtKey = "#$%acd("

// CodeLength 验证码长度
var CodeLength = 6
