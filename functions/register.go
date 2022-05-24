package functions

import (
	"CourseSelectionSystem/db"
	"CourseSelectionSystem/model"
	"errors"
)

// CheckNum 查询用户是否存在
func CheckNum(Num string, role int) error {
	if role == 1 {
		var S model.Student
		db.Mdb.Where("StuNum = ?", Num).First(&S)
		if S.ID > 0 {
			return errors.New("学生学号已存在")
		}
		return nil
	} else if role == 2 {
		var T model.Teacher
		db.Mdb.Where("TeachNum = ?", Num).First(&T)
		if T.ID > 0 {
			return errors.New("教师工号已存在")
		}
		return nil
	} else if role == 3 {
		var A model.Administrator
		db.Mdb.Where("AdminNum = ?", Num).First(&A)
		if A.ID > 0 {
			return errors.New("管理员工号已存在")
		}
		return nil
	}
	return errors.New("角色码错误！")
}
