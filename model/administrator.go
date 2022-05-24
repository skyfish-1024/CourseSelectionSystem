package model

import "github.com/jinzhu/gorm"

//Administrator 管理员
type Administrator struct {
	gorm.Model
	Name     string `gorm:"column:Name;type:varchar(20);not null" json:"Name,omitempty" validate:"required,min=4,max=12"`         //管理员姓名
	Password string `gorm:"column:Password;type:varchar(20);not null" json:"Password,omitempty" validate:"required,min=6,max=20"` //密码
	AdminNum string `gorm:"column:AdminNum;type:varchar(10);not null" json:"AdminNum,omitempty" validate:"required,len=10"`       //工号
	Sex      string `gorm:"column:Sex;type:varchar(2);not null" json:"Sex,omitempty" validate:"required"`                         //性别
	IdCard   string `gorm:"column:IdCard;type:varchar(20);not null" json:"IdCard,omitempty" validate:"required,len=18"`           //身份证
	Phone    string `gorm:"column:Phone;type:varchar(20);not null" json:"Phone,omitempty" validate:"required,len=11"`             //电话
	Email    string `gorm:"column:Email;type:varchar(20);not null" json:"Email,omitempty" validate:"required,len=20"`             //邮箱
	Role     int    `gorm:"column:Role;type:int;DEFAULT:3" json:"Role,omitempty" validate:"required,gte=3" `                      //角色码
}
