package model

import "github.com/jinzhu/gorm"

//Teacher 教师
type Teacher struct {
	gorm.Model
	Name     string `gorm:"column:Name;type:varchar(20);not null" json:"Name,omitempty" validate:"required,max=8"`                //学生姓名
	Password string `gorm:"column:Password;type:varchar(20);not null" json:"Password,omitempty" validate:"required,min=6,max=20"` //密码
	TeachNum string `gorm:"column:TeachNum;type:varchar(10);not null" json:"TeachNum,omitempty" validate:"required,len=10"`       //教师工号
	Grade    int    `gorm:"column:Grade;type:int;DEFAULT:0" json:"Grade,omitempty" validate:"required,max=10"`                    //年级码
	Position string `gorm:"column:Position;type:varchar(20);not null" json:"Position,omitempty" validate:"required,max=20"`       //职位&所任科目
	Sex      string `gorm:"column:Sex;type:varchar(2);not null" json:"Sex,omitempty" validate:"required"`                         //性别
	IdCard   string `gorm:"column:IdCard;type:varchar(20);not null" json:"IdCard,omitempty" validate:"required,len=18"`           //身份证
	Phone    string `gorm:"column:Phone;type:varchar(20);not null" json:"Phone,omitempty" validate:"required,len=11"`             //电话
	Email    string `gorm:"column:Email;type:varchar(20);not null" json:"Email,omitempty" validate:"required,len=20"`             //邮箱
	Role     int    `gorm:"column:Role;type:int;DEFAULT:2" json:"Role,omitempty" validate:"required,gte=3"`                       //角色码
}
