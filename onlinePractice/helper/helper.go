package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s))) //转化为16进制
}

var myKey = "gin-gorm-oj-key"

type UserClaims struct {
	Identity string `json:"identity"`
	Name     string `json:"name"`
	jwt.StandardClaims
}

// 生成token
func GenerateToken(identity, name string) (string, error) {
	UserClaim := &UserClaims{
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString([]byte(myKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// 解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	UserClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, UserClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}

	return UserClaim, nil
}
