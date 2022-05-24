package v1

import (
	"CourseSelectionSystem/functions"
	"CourseSelectionSystem/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//AddCourse 添加课程
func AddCourse(c *gin.Context) {
	var data model.Course
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	err = functions.CheckCourseNum(data.CourseNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	err = functions.CreateCourse(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"status":  "success",
		"message": "创建课程成功",
	})
	return
}

//DeleteCourse 删除课程
func DeleteCourse(c *gin.Context) {
	CourseNum := c.PostForm("CourseNum")
	if CourseNum == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": "课程编号为空删除失败",
		})
		return
	}
	err := functions.DeleteCourse(CourseNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": "删除失败：" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"status":  "success",
		"message": "删除课程成功",
	})
	return
}

//EditeCourse 编辑课程
func EditeCourse(c *gin.Context) {
	var data model.Course
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": "课程信息录入失败",
		})
		return
	}
	err = functions.EditeCourse(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": "更新课程信息失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"status":  "success",
		"message": "更新课程信息成功",
	})
	return
}

//GetCourse 查询单个课程信息
func GetCourse(c *gin.Context) {
	CourseNum := c.PostForm("CourseNum")
	data, err := functions.GetCourse(CourseNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"status":  "success",
		"message": "查询课程信息成功",
		"data:":   data,
	})
}

//GetCourses 查询课程列表
func GetCourses(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	date, total, err := functions.GetCourses(pageSize, pageNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"data":    date,
			"total":   total,
			"message": err.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"status":  "success",
		"data":    date,
		"total":   total,
		"message": "查询成功",
	})
}
