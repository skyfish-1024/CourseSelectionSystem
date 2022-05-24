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

func Login(c *gin.Context) {
	var token string
	Num := c.PostForm("Num")
	Password := c.PostForm("Password")
	Role, _ := strconv.Atoi(c.PostForm("Role"))
	err := functions.CheckLogin(Num, Password, Role)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": err.Error(),
		})
		return
	}
	v, dberr := db.Rdb.HSET("users", Num, Role)
	if dberr != nil && v != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": dberr.Error(),
		})
		return
	}
	//设置token，24小时过期
	token, _ = middleware.SetToken(Num, Role, time.Now().Add(time.Hour*24))
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"status":  "success",
		"message": "验证通过",
		"token":   token,
	})
}

func Logout(c *gin.Context) {
	key, err := middleware.TokenFmtCheck(c)
	if err != nil {
		return
	}
	v, dberr := db.Rdb.HDEL("users", key.Num)
	if dberr != nil && v != 1 {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"status":  "false",
			"message": dberr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"status":  "success",
		"message": "感谢使用！",
	})
	return

}
