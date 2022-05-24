package v1

import (
	"CourseSelectionSystem/functions"
	"CourseSelectionSystem/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//AddStudent 添加学生路由
func AddStudent(c *gin.Context) {
	var data model.Student
	err := c.Bind(&data)
	data.Role = 1
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	err = functions.CheckNum(data.StuNum, data.Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	err = functions.CreateStudent(&data)
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
		"message": "创建学生成功",
	})
	return
}

//DeleteStudent 删除学生路由
func DeleteStudent(c *gin.Context) {
	StuNum := c.PostForm("StuNum")
	if StuNum == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": "学号为空删除失败",
		})
		return
	}
	err := functions.DeleteStudent(StuNum)
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
		"message": "删除学生成功",
	})
	return
}

//EditeStudent 编辑学生路由
func EditeStudent(c *gin.Context) {
	var data model.Student
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": "学生信息录入失败",
		})
		return
	}
	err = functions.EditeStudent(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": "更新学生信息失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    -0,
		"status":  "success",
		"message": "更新学生信息成功",
	})
	return
}

//GetStudent 查询单个学生信息路由
func GetStudent(c *gin.Context) {
	StuNum := c.PostForm("StuNum")
	data, err := functions.GetStudent(StuNum)
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
		"message": "查询学生信息成功",
		"data:":   data,
	})
}

//GetStudents 查询学生列表
func GetStudents(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	date, total, err := functions.GetStudents(pageSize, pageNum)
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
