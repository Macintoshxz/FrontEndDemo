package utils

// 通用工具

import (
	"crypto/md5"
	"errors"
	"fmt"
	"management-system/pkg/define"
	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
)

/**
* @功    能：md5加密
* @输入参数：字符串
* @返 回 值：返回值描述
* @日    期：2023年09月03日11:56:49
* @版    本：版本号
 */
func MD5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

/**
* @功    能：token生成
* @输入参数：输入参数描述
* @返 回 值：返回值描述
* @日    期：2023年09月03日12:04:01
* @版    本：版本号
 */
func GenerateToken(id int, nick, name string, second int) (string, error) {

	uc := define.UserClaim{
		Id:             id,
		Name:           name,
		Nick:           nick,
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Second * time.Duration(second)).Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)

	tokenString, err := token.SignedString([]byte(define.JwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

/**
* @功    能：解析token
* @输入参数：输入参数描述
* @返 回 值：返回值描述
* @日    期：2023年09月03日12:04:23
* @版    本：版本号
 */
func AnalyzeToken(auth string) (*define.UserClaim, error) {
	uc := new(define.UserClaim)

	//拿取clains
	claims, err := jwt.ParseWithClaims(auth, uc, func(t *jwt.Token) (interface{}, error) {
		return []byte(define.JwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return uc, errors.New("token is invalid")
	}
	return uc, err
}

/**
* @功    能：随机0-9验证码
* @输入参数：输入参数描述
* @返 回 值：返回值描述
* @日    期：2023年09月03日12:04:55
* @版    本：版本号
 */
func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < define.CodeLength; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

/**
* @功    能：UUID
* @输入参数：输入参数描述
* @返 回 值：返回值描述
* @日    期：2023年09月03日12:05:43
* @版    本：版本号
 */
func UUID() string {
	return uuid.NewV4().String()
}

/**
* @功    能：生成15位随机字母与数字
* @输入参数：输入参数描述
* @返 回 值：返回值描述
* @日    期：2023年09月14日11:37:29
* @版    本：版本号
 */
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}
