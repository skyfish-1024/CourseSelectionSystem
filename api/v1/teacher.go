package v1

import (
	"CourseSelectionSystem/functions"
	"CourseSelectionSystem/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//AddTeacher 添加教师路由
func AddTeacher(c *gin.Context) {
	var data model.Teacher
	err := c.Bind(&data)
	data.Role = 2
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "False",
			"message": err.Error(),
		})
		return
	}
	err = functions.CheckNum(data.TeachNum, data.Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "False",
			"message": err.Error(),
		})
		return
	}
	err = functions.CreateTeacher(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "False",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "创建教师成功",
	})
	return
}

//DeleteTeacher 删除教师路由
func DeleteTeacher(c *gin.Context) {
	TeachNum := c.PostForm("TeachNum")
	if TeachNum == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  "False",
			"message": "工号为空删除失败",
		})
		return
	}
	err := functions.DeleteTeacher(TeachNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "False",
			"message": "删除失败：" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "删除教师成功",
	})
	return
}

//EditeTeacher 编辑教师信息路由
func EditeTeacher(c *gin.Context) {
	var data model.Teacher
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "FALSE",
			"message": "学生信息录入失败",
		})
		return
	}
	data.Password = ""
	err = functions.EditeTeacher(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "FALSE",
			"message": "更新教师信息失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "更新教师信息成功",
	})
	return
}

//GetTeacher 查询单个教师信息路由
func GetTeacher(c *gin.Context) {
	TeachNum := c.PostForm("TeachNum")
	data, err := functions.GetTeacher(TeachNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "FALSE",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "查询教师信息成功",
		"data:":   data,
	})
}

//GetTeachers 查询教师列表路由
func GetTeachers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	date, total, err := functions.GetTeachers(pageSize, pageNum)
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
