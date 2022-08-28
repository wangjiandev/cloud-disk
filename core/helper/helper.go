package helper

import (
	"context"
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"path"
	"time"
	"wangjiandev/cloud-disk/core/define"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jordan-wright/email"
	uuid "github.com/satori/go.uuid"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// 生成MD5
func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// 生成token
func GenerateToken(id uint64, identity string, name string) (string, error) {
	uc := define.UserClaim{}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
	signedString, err := token.SignedString([]byte(define.Jwtkey))
	if err != nil {
		return "", err
	}
	return signedString, nil
}

func EmailSendCode(mailAddress string, code string) error {
	e := email.NewEmail()
	e.From = "Get <wangjiantx2@163.com>"
	e.To = []string{mailAddress}
	e.Subject = "验证码"
	e.HTML = []byte("<h1>您的验证码: " + code + " </h1>")
	err := e.SendWithTLS("smtp.163.com:465", smtp.PlainAuth("", "wangjiantx2@163.com", define.MailKey, "smtp.163.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtp.163.com"})
	if err != nil {
		return err
	}
	fmt.Println("发送成功!")
	return nil
}

func RandCode() string {
	s := "1234567890"
	code := ""
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		code += string(s[rand.Intn(len(s))])
	}
	return code
}

func RandUUID() string {
	s := uuid.NewV4()
	return s.String()
}

func CosUpload(r *http.Request) (string, error) {
	u, _ := url.Parse(define.CosBucketDomain)
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  define.SecretID,
			SecretKey: define.SecretKey,
		},
	})

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	key := "cloud-disk/" + RandUUID() + path.Ext(fileHeader.Filename)
	// 通过文件流上传对象
	_, err = c.Object.Put(context.Background(), key, file, nil)
	if err != nil {
		panic(err)
	}
	return define.CosBucketDomain + "/" + key, nil
}
