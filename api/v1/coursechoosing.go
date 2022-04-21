package v1

import (
	"CourseSelectionSystem/functions"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//CourseChoose 选课路由
func CourseChoose(c *gin.Context) {
	CourseNum := c.PostForm("CourseNum")
	StuNUm := c.PostForm("StuNum")
	err := functions.CourseChoose(CourseNum, StuNUm)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"stata":   "FALSE",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"stata":   "ok",
		"message": "选课成功",
	})
}

//CourseDelStudent 选课结束后管理员删除课程学生
func CourseDelStudent(c *gin.Context) {
	CourseNum := c.PostForm("CourseNum")
	StuNUm := c.PostForm("StuNum")
	err := functions.CourseDelStudent(CourseNum, StuNUm)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"stata":   "FALSE",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"stata":   "ok",
		"message": "退课成功",
	})
}

//CourseAddStudent 选课结束后管理员添加课程学生
func CourseAddStudent(c *gin.Context) {
	CourseNum := c.PostForm("CourseNum")
	StuNUm := c.PostForm("StuNum")
	err := functions.CourseAddStudent(CourseNum, StuNUm)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"stata":   "FALSE",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"stata":   "ok",
		"message": "添加成功",
	})
}

//CourseDrop 退课路由
func CourseDrop(c *gin.Context) {
	CourseNum := c.PostForm("CourseNum")
	StuNUm := c.PostForm("StuNum")
	err := functions.CourseDrop(CourseNum, StuNUm)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"stata":   "FALSE",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"stata":   "ok",
		"message": "退课成功",
	})
}

//GetChoseCourses 查询学生已选课程列表
func GetChoseCourses(c *gin.Context) {
	StuNum := c.PostForm("StuNum")
	date, total, err := functions.GetChoseCourses(StuNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "FALSE",
			"data":    date,
			"total":   total,
			"message": err.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"data":    date,
		"total":   total,
		"message": "查询成功",
	})
}

//GetCourseChoseStu 查询课程列表
func GetCourseChoseStu(c *gin.Context) {
	CourseNum := c.PostForm("CourseNum")
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	date, total, err := functions.GetCourseChoseStu(CourseNum, pageSize, pageNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "FALSE",
			"data":    date,
			"total":   total,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"data":    date,
		"total":   total,
		"message": "查询成功",
	})
}

//OpenCourse 开放课程选课
func OpenCourse(c *gin.Context) {
	CourseNum := c.PostForm("CourseNum")
	err := functions.OpenCourse(CourseNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "FALSE",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": CourseNum + "课程开放成功",
	})
	return
}

//CloseCourse 关闭课程选课
func CloseCourse(c *gin.Context) {
	CourseNum := c.PostForm("CourseNum")
	errStu, err := functions.CloseCourse(CourseNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":         "FALSE",
			"message":        err.Error(),
			"error students": errStu,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":         "success",
		"message":        CourseNum + "课程关闭成功",
		"error students": errStu,
	})
	return
}
