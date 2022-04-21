package v1

import (
	"CourseSelectionSystem/db"
	"CourseSelectionSystem/functions"
	"CourseSelectionSystem/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// SendEmail 邮件发送路由
func SendEmail(c *gin.Context) {
	Num := c.PostForm("Num")
	UserMail := c.PostForm("UserMail")
	Role, _ := strconv.Atoi(c.PostForm("Role"))
	err := functions.MailCheck(Num, UserMail, Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "False",
			"message": err.Error(),
		})
		return
	}
	VC, Merr := functions.SendMail(UserMail, Num)
	if Merr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "False",
			"message": err.Error(),
		})
		return
	}
	Rdb := db.NewClient()
	err = Rdb.VCset(Num, VC)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "False",
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "验证码已发送",
	})
	return
}

// VCCheck 验证码校验路由
func VCCheck(c *gin.Context) {
	Num := c.PostForm("Num")
	UserMail := c.PostForm("UserMail")
	Role, _ := strconv.Atoi(c.PostForm("Role"))
	VC := c.PostForm("VC")
	err := functions.MailCheck(Num, UserMail, Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "False",
			"message": err.Error(),
		})
		return
	}
	err = functions.VCCheck(Num, VC)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "False",
			"message": err.Error(),
		})
		return
	}

	Role = 68 //将角色码改为68，作为重置密码的权限验证
	token, _ := middleware.SetToken(Num, Role, time.Now().Add(time.Second*600))
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "验证通过",
		"token":   token,
	})
	return
}

func PasswordReset(c *gin.Context) {
	Num := c.Query("Num")
	Role, _ := strconv.Atoi(c.Query("Role"))
	Password := c.PostForm("Password")
	err := functions.PasswordReset(Num, Password, Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "密码修改成功",
	})
	return
}
