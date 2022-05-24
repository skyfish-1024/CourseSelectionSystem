package v1

import (
	"CourseSelectionSystem/db"
	"CourseSelectionSystem/functions"
	"CourseSelectionSystem/middleware"
	"CourseSelectionSystem/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func RegisterMail(c *gin.Context) {
	Num := c.PostForm("Num")
	UserMail := c.PostForm("UserMail")
	Role, _ := strconv.Atoi(c.PostForm("Role"))
	err := functions.CheckNum(Num, Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	VC, Merr := functions.SendMail(UserMail, Num, 3)
	if Merr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	err = db.Rdb.VCset(Num, VC)
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
		"message": "验证码已发送",
	})
	return
}

// VCCheck 验证码校验路由
func RVCCheck(c *gin.Context) {
	Num := c.PostForm("Num")
	Role, _ := strconv.Atoi(c.PostForm("Role"))
	VC := c.PostForm("VC")
	err := functions.VCCheck(Num, VC)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	token, _ := middleware.SetToken(Num, Role, time.Now().Add(time.Second*600))
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"status":  "success",
		"message": "验证通过",
		"token":   token,
	})
	return
}

func UserInfo(c *gin.Context) {
	var data model.Administrator
	key, err := middleware.TokenFmtCheck(c)
	if err != nil {
		return
	}
	err = c.Bind(&data)
	data.AdminNum = key.Num
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
