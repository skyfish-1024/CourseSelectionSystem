package functions

import (
	"CourseSelectionSystem/db"
	"CourseSelectionSystem/model"
	"errors"
	"github.com/jinzhu/gorm"
)

// CreateTeacher 新增教师
func CreateTeacher(data *model.Teacher) error {
	data.Password = ScryptPw(data.Password) //密码加密
	err := db.Mdb.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

// EditeTeacher 编辑教师信息
func EditeTeacher(data *model.Teacher) error {
	data.Password = ScryptPw(data.Password)
	//maps := make(map[string]interface{})
	//maps["Name"] = data.Name
	//maps["Password"] = data.Password
	//maps["TeachNum"] = data.TeachNum
	//maps["Grade"] = data.Grade
	//maps["Position"] = data.Position
	//maps["Sex"] = data.Sex
	//maps["IdCard"] = data.IdCard
	//maps["Phone"] = data.Phone
	//maps["Email"] = data.Email
	//maps["Role"] = data.Role
	err := db.Mdb.Model(&model.Teacher{}).Where("TeachNum=?", data.TeachNum).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteTeacher 删除教师
func DeleteTeacher(TeachNum string) error {
	err := db.Mdb.Where("TeachNum=?", TeachNum).Delete(&model.Teacher{}).Error
	if err != nil {
		return err
	}
	_, err = db.Rdb.HDEL("users", TeachNum)
	if err != nil {
		return err
	}
	return nil
}

// GetTeachers 查询教师列表
func GetTeachers(pageSize int, pageNum int) ([]model.Teacher, int, error) {
	var Teachers []model.Teacher
	var total int
	err := db.Mdb.Select("Name,TeachNum,Grade,Position,Sex,Phone,Email").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&Teachers).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return Teachers, total, nil
}

//GetTeacher 查询单个教师信息
func GetTeacher(TeachNum string) (model.Teacher, error) {
	var teacher model.Teacher
	err := db.Mdb.Select("Name,TeachNum,Grade,Position,Sex,Phone,Email").Where("TeachNum=?", TeachNum).First(&teacher).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return teacher, errors.New("工号不存在")
		}
		return teacher, nil
	}
	return teacher, nil
}
