package functions

import (
	"CourseSelectionSystem/db"
	"CourseSelectionSystem/model"
	"errors"
	"github.com/jinzhu/gorm"
)

// CreateCourse 新增课程
func CreateCourse(data *model.Course) error {
	err := db.Mdb.Create(data).Error
	if err != nil {
		return err
	}
	return nil
}

// EditeCourse 编辑课程信息
func EditeCourse(data *model.Course) error {
	//maps := make(map[string]interface{})
	//maps["CourseName"] = data.CourseName
	//maps["CourseNum"] = data.CourseNum
	//maps["Grade"] = data.Grade
	//maps["Period"] = data.Period
	//maps["Score"] = data.Score
	//maps["Teacher"] = data.Teacher
	//maps["TeachNum"] = data.TeachNum
	//maps["TotalStu"] = data.TotalStu
	//maps["LeftStu"] = data.LeftStu
	//maps["Time"] = data.Time
	//maps["Description"] = data.Description
	//maps["State"] = data.State
	err := db.Mdb.Model(&model.Course{}).Where("CourseNum=?", data.CourseNum).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

// DeleteCourse 删除课程
func DeleteCourse(CourseNum string) error {
	err := db.Mdb.Where("CourseNum=?", CourseNum).Delete(&model.Course{}).Error
	if err != nil {
		return err
	}
	return nil
}

// GetCourses 查询课程列表
func GetCourses(pageSize int, pageNum int) ([]model.Course, int, error) {
	var Courses []model.Course
	var total int
	err := db.Mdb.Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&Courses).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return Courses, total, nil
}

//GetCourse 查询单个课程信息
func GetCourse(CourseNum string) (model.Course, error) {
	var Course model.Course
	err := db.Mdb.Where("CourseNum=?", CourseNum).First(&Course).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return Course, errors.New("课程号不存在")
		}
		return Course, err
	}
	return Course, nil
}

//CheckCourseNum 查验课程是否存在
func CheckCourseNum(CourseNum string) error {
	var C model.Course
	db.Mdb.Where("CourseNum = ?", CourseNum).First(&C)
	if C.ID > 0 {
		return errors.New("课程已存在")
	}
	return nil
}
