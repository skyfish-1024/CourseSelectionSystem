package model

import (
	"github.com/jinzhu/gorm"
)

//CourseChoosing 已选课程
type CourseChoosing struct {
	gorm.Model
	CourseName string `gorm:"column:CourseName;type:varchar(20);not null" json:"CourseName,omitempty" validate:"required,max=12"` //课程名称
	CourseNum  string `gorm:"column:CourseNum;type:varchar(10);not null" json:"CourseNum,omitempty" validate:"required,max=10"`   //课程码
	Teacher    string `gorm:"column:Teacher;type:varchar(20);not null" json:"Teacher,omitempty" validate:"required,max=12"`       //任课老师
	StuName    string `gorm:"column:StuName;type:varchar(20);not null;" json:"Name,omitempty" validate:"required,min=4,max=12"`   //学生姓名
	StuNum     string `gorm:"column:StuNum;varchar(10);not null;" json:"StuNum,omitempty" validate:"required,len=10"`             //学号
	Day        int    `gorm:"column:Day;type:int;DEFAULT:0" json:"Day,omitempty" validate:"required,max=10" `                     //星期
	Time       int    `gorm:"column:Time;type:int;DEFAULT:0" json:"Time,omitempty" validate:"required,max=10"`                    //时间
}
