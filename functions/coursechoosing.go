package functions

import (
	"CourseSelectionSystem/db"
	"CourseSelectionSystem/model"
	"errors"
	"github.com/jinzhu/gorm"
)

//CourseAddStudent MySQL中向课程中添加学生
func CourseAddStudent(CourseNum string, StuNUm string) error {
	var course model.Course
	var student model.Student
	var courseChoose model.CourseChoosing
	err := db.Mdb.Where("StuNUm=?", StuNUm).First(&student).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("学号不存在")
		}
		return err
	}
	err = db.Mdb.Where("CourseNum=? AND StuNUm=?", CourseNum, StuNUm).First(&courseChoose).Error
	if err != gorm.ErrRecordNotFound {
		return errors.New("你已选择该课程")
	}
	err = db.Mdb.Where("CourseNum=?", CourseNum).First(&course).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("课程不存在")
		}
		return err
	}
	if student.Grade != course.Grade {
		return errors.New("不能跨年级选课")
	}
	//else if course.State == 2 {
	//	return errors.New("该课程未开放")
	//}
	maps := make(map[string]interface{})
	maps["LeftStu"] = course.LeftStu - 1
	err = db.Mdb.Model(&model.Course{}).Where("CourseNum=?", course.CourseNum).Updates(maps).Error
	if err != nil {
		return errors.New("选课失败")
	}
	courseChoose.CourseName = course.CourseName
	courseChoose.CourseNum = course.CourseNum
	courseChoose.Teacher = course.Teacher
	courseChoose.StuName = student.Name
	courseChoose.StuNum = student.StuNum
	courseChoose.Day = course.Day
	courseChoose.Time = course.Time
	err = db.Mdb.Create(&courseChoose).Error
	if err != nil {
		return errors.New("选课失败")
	}

	return nil
}

//CourseDelStudent MySQL中向课程中删除学生
func CourseDelStudent(CourseNum string, StuNUm string) error {
	var course model.Course
	err := db.Mdb.Where("CourseNum=? AND StuNUm=?", CourseNum, StuNUm).First(&model.CourseChoosing{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("未选择该课程")
		}
		return err
	}
	err = db.Mdb.Where("CourseNum=?", CourseNum).First(&course).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("课程不存在，无法操作")
		}
		return err
	}
	if course.State == 1 {
		return errors.New("该课程选课还未结束，无法操作，删除失败")
	}
	err = db.Mdb.Where("CourseNum=?", CourseNum).Where("StuNUm=?", StuNUm).Unscoped().Delete(&model.CourseChoosing{}).Error
	if err != nil {
		return errors.New("删除失败")
	}
	maps := make(map[string]interface{})
	maps["LeftStu"] = course.LeftStu + 1
	err = db.Mdb.Model(&model.Course{}).Where("CourseNum=?", course.CourseNum).Updates(maps).Error
	if err != nil {
		return errors.New("删除失败")
	}
	return nil
}

//GetChoseCourses 学生查询已选课程
//func GetChoseCourses(StuNum string) ([]model.CourseChoosing, int, error) {
//	var Courses []model.CourseChoosing
//	var total int
//	err := db.Mdb.Where("StuNum=?", StuNum).Find(&Courses).Count(&total).Error
//	if err != nil && err != gorm.ErrRecordNotFound {
//		return nil, 0, err
//	}
//	return Courses, total, nil
//}

func GetChoseCourses(StuNum string) ([]string, int, error) {
	Rdb := db.NewClient()
	Courses, err := Rdb.GetStuCourse(StuNum)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return Courses, len(Courses), nil
}

//GetCourseChoseStu 教师查询已选课程学生
func GetCourseChoseStu(CourseNum string, pageSize int, pageNum int) ([]model.CourseChoosing, int, error) {
	var Courses []model.CourseChoosing
	var Course model.Course
	var total int
	err := db.Mdb.Where("CourseNum=?", CourseNum).First(&Course).Error
	if err != nil {
		return Courses, 0, err
	}
	if Course.State != 2 {
		return Courses, 0, errors.New("该课程选课还未结束，无法查询")
	}
	err = db.Mdb.Where("CourseNum=?", CourseNum).Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&Courses).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return Courses, 0, err
	}
	return Courses, total, nil
}

//OpenCourse 发布可选课程
func OpenCourse(CourseNum string) error {
	var data model.Course
	err := db.Mdb.Where("CourseNum=?", CourseNum).First(&data).Error
	if err != nil {
		return err
	}
	if data.State == 1 {
		return errors.New("该课程已开放")
	}
	LeftStu := data.LeftStu
	data = model.Course{}
	data.State = 1 //改变课程状态为可变
	data.CourseNum = CourseNum
	err = EditeCourse(&data)
	if err != nil {
		return err
	}
	Rdb := db.NewClient()
	_, err = Rdb.OpenCourse(CourseNum, LeftStu)
	if err != nil {
		return err
	}
	return nil
}

//CloseCourse 关闭可选课程
func CloseCourse(CourseNum string) ([]string, error) {
	var data model.Course
	var errStu []string
	err := db.Mdb.Where("CourseNum=?", CourseNum).First(&data).Error
	if err != nil {
		return errStu, err
	}
	if data.State == 2 {
		return errStu, errors.New("该课程已关闭")
	}
	data = model.Course{}
	data.State = 2 //改变课程状态为不可变
	data.CourseNum = CourseNum
	err = EditeCourse(&data)
	if err != nil {
		return errStu, err
	}
	Rdb := db.NewClient()
	students, err1 := Rdb.CloseCourse(CourseNum)
	if err != nil {
		return errStu, err1
	}
	for i := 0; i < len(students); i++ {
		err = CourseAddStudent(CourseNum, students[i])
		if err != nil && err.Error() != "你已选择该课程" {
			errStu = append(errStu, students[i]+":"+err.Error())
		}
	}
	return errStu, nil
}

//CourseChoose 选课
func CourseChoose(CourseNum string, StuNUm string) error {
	Rdb := db.NewClient()
	err := Rdb.CourseAddStu(CourseNum, StuNUm)
	if err != nil {
		return err
	}
	return nil
}

//CourseDrop 退课
func CourseDrop(CourseNum string, StuNUm string) error {
	Rdb := db.NewClient()
	v, err := Rdb.CourseDelStu(CourseNum, StuNUm)
	if err != nil {
		return err
	}
	if v.(int64) != 1 {
		return errors.New("退课失败")
	}
	return nil
}
