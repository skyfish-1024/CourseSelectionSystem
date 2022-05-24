package functions

import (
	"CourseSelectionSystem/db"
	"CourseSelectionSystem/model"
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/scrypt"
	"log"
)

//ScryptPw 密码加密
func ScryptPw(password string) string {
	const KeyLen = 10
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 6, 66, 22, 222, 11}
	HashPw, err := scrypt.Key([]byte(password), salt, 16384, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(HashPw)
	return fpw
}

// CheckLogin 登录验证
func CheckLogin(Num string, password string, role int) error {
	if role == 1 {
		var S model.Student
		db.Mdb.Where("StuNum=?", Num).First(&S)
		if S.ID == 0 {
			err := errors.New("学号不存在")
			return err
		}
		if ScryptPw(password) != S.Password {
			return errors.New("密码错误")
		}
	} else if role == 2 {
		var T model.Teacher
		db.Mdb.Where("TeachNum=?", Num).First(&T)
		if T.ID == 0 {
			return errors.New("工号不存在")
		}
		if ScryptPw(password) != T.Password {
			return errors.New("密码错误")
		}
	} else if role == 3 {
		var A model.Administrator
		db.Mdb.Where("AdminNum=?", Num).First(&A)
		if A.ID == 0 {
			return errors.New("工号不存在")
		}
		if ScryptPw(password) != A.Password {
			return errors.New("密码错误")
		}
	} else {
		return errors.New("角色错误！")
	}
	return nil
}
