package functions

import (
	"CourseSelectionSystem/config"
	"errors"
	"github.com/jordan-wright/email"
	"log"
	"math/rand"
	"net/smtp"
	"time"
)

func SendMail(UserEmail string, Num string, code int) (string, error) {
	// 简单设置 log 参数
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "教务处 <" + config.FromEmail + ">"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{UserEmail}

	// 设置主题
	em.Subject = "验证码"

	VC := RandomString(5)
	// 简单设置文件发送的内容，暂时设置成纯文本
	if code == 6 {
		//code=6，发送找回密码验证信息
		em.Text = []byte("【教务处】尊敬的 " + Num + " 用户，您正在通过手机邮箱更改登陆密码，您的验证码为：" + VC + "，该验证码 10 分钟内有效，请勿泄漏于他人。\n")
	} else if code == 3 {
		//code=3，发送管理员注册信息
		em.Text = []byte("【教务处】尊敬的 " + Num + " 管理员用户，您正在通过手机邮箱注册账号，您的验证码为：" + VC + "，该验证码 10 分钟内有效，请勿泄漏于他人。\n")
	} else {
		return "", errors.New("邮件类型有误")
	}

	//设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "3574819459@qq.com", config.EVC, "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
		return VC, errors.New("验证码发送失败！")
	}
	log.Println("send successfully ... ")
	return VC, nil
}

//RandomString 生成随机字符串，作为验证码
func RandomString(len int) string {
	bytes := make([]byte, len)
	rand.Seed(time.Now().Unix())
	for i := 0; i < len; i++ {
		bytes[i] = byte(49 + rand.Intn(9))
	}
	return string(bytes)
}
