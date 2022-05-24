package functions

import (
	"CourseSelectionSystem/db"
	"CourseSelectionSystem/model"
	"errors"
)

// PasswordReset 密码重置
func PasswordReset(Num string, Password string, Role int) (err error) {
	if Role == 1 {
		data := model.Student{}
		data.StuNum = Num
		data.Password = Password
		err = EditeStudent(&data)
		if err != nil {
			return err
		}
	} else if Role == 2 {
		data := model.Teacher{}
		data.TeachNum = Num
		data.Password = Password
		err = EditeTeacher(&data)
		if err != nil {
			return err
		}
	} else if Role == 3 {
		data := model.Administrator{}
		data.AdminNum = Num
		data.Password = Password
		err = EditeAdministrator(&data)
		if err != nil {
			return err
		}
	} else {
		return errors.New("角色码错误")
	}
	return nil
}

// VCCheck 验证码校验
func VCCheck(Num string, VC string) error {
	rvc, err := db.Rdb.VCget(Num)
	if err != nil {
		return err
	}
	if rvc != VC {
		return errors.New("验证码错误")
	}
	return nil
}

// MailCheck 邮箱校验
func MailCheck(Num string, UserMail string, role int) error {
	if role == 1 {
		var S model.Student
		db.Mdb.Where("StuNum=?", Num).First(&S)
		if S.ID == 0 {
			err := errors.New("学号不存在")
			return err
		}
		if UserMail != S.Email {
			return errors.New("学号和邮箱不匹配")
		}
	} else if role == 2 {
		var T model.Teacher
		db.Mdb.Where("TeachNum=?", Num).First(&T)
		if T.ID == 0 {
			return errors.New("工号不存在")
		}
		if UserMail != T.Email {
			return errors.New("工号和邮箱不匹配")
		}
	} else if role == 3 {
		var A model.Administrator
		db.Mdb.Where("AdminNum=?", Num).First(&A)
		if A.ID == 0 {
			return errors.New("工号不存在")
		}
		if UserMail != A.Email {
			return errors.New("工号和邮箱不匹配")
		}
	} else {
		return errors.New("角色错误！")
	}
	return nil
}
