package functions

import (
	"CourseSelectionSystem/db"
	"CourseSelectionSystem/model"
	"errors"
	"github.com/jinzhu/gorm"
)

// CreateStudent 新增学生
func CreateStudent(data *model.Student) error {
	data.Password = ScryptPw(data.Password) //密码加密
	err := db.Mdb.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

// EditeStudent 编辑学生信息
func EditeStudent(data *model.Student) error {
	data.Password = ScryptPw(data.Password)
	//maps := make(map[string]interface{})
	//maps["Name"] = data.Name
	//maps["Password"] = ScryptPw(data.Password)
	//maps["StuNum"] = data.StuNum
	//maps["Grade"] = data.Grade
	//maps["Class"] = data.Class
	//maps["Major"] = data.Major
	//maps["Sex"] = data.Sex
	//maps["IdCard"] = data.IdCard
	//maps["Phone"] = data.Phone
	//maps["Email"] = data.Email
	err := db.Mdb.Model(&model.Student{}).Where("StuNum=?", data.StuNum).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteStudent 删除学生
func DeleteStudent(StuNum string) error {
	err := db.Mdb.Where("StuNum=?", StuNum).Delete(&model.Student{}).Error
	if err != nil {
		return err
	}
	err = db.Mdb.Where("StuNum=?", StuNum).Delete(&model.CourseChoosing{}).Error
	if err != nil {
		return err
	}
	_, err = db.Rdb.HDEL("users", StuNum)
	if err != nil {
		return err
	}
	return nil
}

//GetStudent 查询单个学生信息
func GetStudent(StuNum string) (model.Student, error) {
	var student model.Student
	err := db.Mdb.Select("Name,StuNum,Grade,Class,Major,Sex,Phone,Email").Where("StuNum=?", StuNum).First(&student).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return student, errors.New("学号不存在")
		}
		return student, nil
	}
	return student, nil
}

// GetStudents 查询学生列表,分页功能
func GetStudents(pageSize int, pageNum int) ([]model.Student, int, error) {
	var Students []model.Student
	var total int
	err := db.Mdb.Select("Name,StuNum,Grade,Class,Major,Sex,Phone,Email").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&Students).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return Students, total, nil
}
