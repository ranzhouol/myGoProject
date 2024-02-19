package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"net/smtp"
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

// SendCode 发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "测试人 <ranzhouol@163.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码发送测试"
	e.Text = []byte("Text Body is, of course, supported!")
	e.HTML = []byte("你的验证码是:<b>" + code + "</b>")
	// 返回 EOF 错误时，关闭SSL
	err := e.SendWithTLS("smtp.163.com:465",
		smtp.PlainAuth("", "ranzhouol@163.com", "QFTPLLCJVEWAVEUX", "smtp.163.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}

	return nil
}
