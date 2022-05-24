package model

import (
	"github.com/jinzhu/gorm"
)

//Course 课程
type Course struct {
	gorm.Model
	CourseName  string `gorm:"column:CourseName;type:varchar(20);not null" json:"CourseName,omitempty" validate:"required,max=12"`   //课程名称
	CourseNum   string `gorm:"column:CourseNum;type:varchar(10);not null" json:"CourseNum,omitempty" validate:"required,max=10"`     //课程码
	Grade       int    `gorm:"column:Grade;type:int;DEFAULT:0" json:"Grade,omitempty" validate:"required,max=10"`                    //开设年级码
	Period      int    `gorm:"column:Period;type:int;DEFAULT:0" json:"Period,omitempty" validate:"required,max=10" `                 //学时
	Score       int    `gorm:"column:Score;type:int;DEFAULT:0" json:"Score,omitempty" validate:"required,max=10"`                    //学分
	Teacher     string `gorm:"column:Teacher;type:varchar(20);not null" json:"Teacher,omitempty" validate:"required,max=12"`         //任课老师
	TeachNum    string `gorm:"column:TeachNum;type:varchar(10);not null" json:"TeachNum,omitempty" validate:"required,len=10"`       //教师工号
	TotalStu    int    `gorm:"column:TotalStu;type:int;DEFAULT:0" json:"TotalStu,omitempty" validate:"required,max=10"`              //学生总人数
	LeftStu     int    `gorm:"column:LeftStu;type:int;DEFAULT:0" json:"LeftStu,omitempty" validate:"required,max=10" `               //剩余可选人数
	Day         int    `gorm:"column:Day;type:int;DEFAULT:0" json:"Day,omitempty" validate:"required,max=10" `                       //星期
	Time        int    `gorm:"column:Time;type:int;DEFAULT:0" json:"Time,omitempty" validate:"required,max=10" `                     //时间
	Description string `gorm:"column:Description;type:varchar(20);not null" json:"Description,omitempty" validate:"required,max=12"` //课程描述
	State       int    `gorm:"column:State;type:int;DEFAULT:2" json:"State,omitempty" validate:"required,max=10" `                   //课程状态:1可变，2不可变
}
