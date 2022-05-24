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
func BackEmail(c *gin.Context) {
	Num := c.PostForm("Num")
	UserMail := c.PostForm("UserMail")
	Role, _ := strconv.Atoi(c.PostForm("Role"))
	err := functions.MailCheck(Num, UserMail, Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	VC, Merr := functions.SendMail(UserMail, Num, 6)
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
func PVCCheck(c *gin.Context) {
	Num := c.PostForm("Num")
	UserMail := c.PostForm("UserMail")
	Role, _ := strconv.Atoi(c.PostForm("Role"))
	VC := c.PostForm("VC")
	err := functions.MailCheck(Num, UserMail, Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	err = functions.VCCheck(Num, VC)
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

// PasswordReset 重置密码路由
func PasswordReset(c *gin.Context) {
	key, err := middleware.TokenFmtCheck(c)
	Password := c.PostForm("Password")
	err = functions.PasswordReset(key.Num, Password, key.Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"status":  "success",
		"message": "密码修改成功",
	})
	return
}
