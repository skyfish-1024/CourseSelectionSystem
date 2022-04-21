package model

import "github.com/jinzhu/gorm"

//Student 学生
type Student struct {
	gorm.Model
	Name     string `gorm:"column:Name;type:varchar(20);not null;" json:"Name,omitempty" validate:"required,min=4,max=12" `        //学生姓名
	Password string `gorm:"column:Password;type:varchar(20);not null;" json:"Password,omitempty" validate:"required,min=6,max=20"` //密码
	StuNum   string `gorm:"column:StuNum;varchar(10);not null;" json:"StuNum,omitempty" validate:"required,len=10"`                //学号
	Grade    int    `gorm:"column:Grade;type:int;DEFAULT:0;" json:"Grade,omitempty" validate:"required"`                           //年级码
	Class    int    `gorm:"column:Class;type:int;DEFAULT:0;" json:"Class,omitempty" validate:"required"`                           //班级码
	Major    string `gorm:"column:Major;type:varchar(20);not null;" json:"Major,omitempty" validate:"required,min=2,max=20"`       //专业
	Sex      string `gorm:"column:Sex;type:varchar(2);not null;" json:"Sex,omitempty" validate:"required"`                         //性别
	IdCard   string `gorm:"column:IdCard;type:varchar(20);not null;" json:"IdCard,omitempty" validate:"required,len=18"`           //身份证
	Phone    string `gorm:"column:Phone;type:varchar(20);not null;" json:"Phone,omitempty" validate:"required,len=11"`             //电话
	Email    string `gorm:"column:Email;type:varchar(20);not null" json:"Email,omitempty" validate:"required,len=20"`              //邮箱
	Role     int    `gorm:"column:Role;type:int;DEFAULT:1;" json:"Role,omitempty" validate:"required,gte=3"`                       //角色码
}
