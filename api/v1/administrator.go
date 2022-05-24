package v1

import (
	"CourseSelectionSystem/functions"
	"CourseSelectionSystem/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//AddAdministrator 添加管理员
func AddAdministrator(c *gin.Context) {
	var data model.Administrator
	err := c.Bind(&data)
	data.Role = 3
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	err = functions.CheckNum(data.AdminNum, data.Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	err = functions.CreateAdministrator(&data)
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
		"message": "创建管理员成功",
	})
	return
}

//DeleteAdministrator 删除管理员
func DeleteAdministrator(c *gin.Context) {
	AdminNum := c.PostForm("AdminNum")
	if AdminNum == "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": "工号为空删除失败",
		})
		return
	}
	err := functions.DeleteAdministrator(AdminNum)
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
		"message": "删除管理员成功",
	})
	return
}

//EditeAdministrator 编辑管理员
func EditeAdministrator(c *gin.Context) {
	var data model.Administrator
	err := c.Bind(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": "管理员信息录入失败",
		})
		return
	}
	data.Password = ""
	err = functions.EditeAdministrator(&data)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": "更新管理员信息失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"status":  "success",
		"message": "更新管理员信息成功",
	})
	return
}

//GetAdministrator 查询单个管理员信息
func GetAdministrator(c *gin.Context) {
	AdminNum := c.PostForm("AdminNum")
	data, err := functions.GetAdministrator(AdminNum)
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
		"message": "查询管理员信息成功",
		"data:":   data,
	})
}

//GetAdministrators 查询管理员列表
func GetAdministrators(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	date, total, err := functions.GetAdministrators(pageSize, pageNum)
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
