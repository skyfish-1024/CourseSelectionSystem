package functions

import (
	"CourseSelectionSystem/db"
	"CourseSelectionSystem/model"
	"errors"
	"github.com/jinzhu/gorm"
)

// CreateAdministrator 新增管理员
func CreateAdministrator(data *model.Administrator) error {
	data.Password = ScryptPw(data.Password) //密码加密
	err := db.Mdb.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

// EditeAdministrator 编辑管理员信息
func EditeAdministrator(data *model.Administrator) error {
	data.Password = ScryptPw(data.Password)
	//maps := make(map[string]interface{})
	//maps["Name"] = data.Name
	//maps["Password"] = data.Password
	//maps["AdminNum"] = data.AdminNum
	//maps["Sex"] = data.Sex
	//maps["IdCard"] = data.IdCard
	//maps["Phone"] = data.Phone
	//maps["Email"] = data.Email
	//maps["Role"] = data.Role
	err := db.Mdb.Model(&model.Administrator{}).Where("AdminNum=?", data.AdminNum).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteAdministrator 删除管理员
func DeleteAdministrator(AdminNum string) error {
	err := db.Mdb.Where("AdminNum=?", AdminNum).Delete(&model.Administrator{}).Error
	if err != nil {
		return err
	}
	Rdb := db.NewClient()
	_, err = Rdb.HDEL("users", AdminNum)
	if err != nil {
		return err
	}
	return nil
}

// GetAdministrators 查询管理员列表
func GetAdministrators(pageSize int, pageNum int) ([]model.Administrator, int, error) {
	var Administrators []model.Administrator
	var total int
	err := db.Mdb.Select("Name,AdminNum,Sex,Phone,Email").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&Administrators).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return Administrators, total, nil
}

//GetAdministrator 查询单个管理员信息
func GetAdministrator(AdminNum string) (model.Administrator, error) {
	var administrator model.Administrator
	err := db.Mdb.Select("Name,AdminNum,Sex,Phone,Email").Where("AdminNum=?", AdminNum).First(&administrator).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return administrator, errors.New("工号不存在")
		}
		return administrator, nil
	}
	return administrator, nil
}
