package v1

import (
	"CourseSelectionSystem/db"
	"CourseSelectionSystem/functions"
	"CourseSelectionSystem/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
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
			"status":  "False",
			"message": err.Error(),
		})
		return
	}
	Rdb := db.NewClient()
	v, dberr := Rdb.HSET("users", Num, Role)
	if dberr != nil && v != 1 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "False",
			"message": dberr.Error(),
		})
		return
	}
	//设置token，24小时过期
	token, _ = middleware.SetToken(Num, Role, time.Now().Add(time.Hour*24))
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "验证通过",
		"token":   token,
	})
}

func Logout(c *gin.Context) {
	tokenHeader := c.Request.Header.Get("Authorization")
	if tokenHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"stata":   "FALSE",
			"message": "token不存在，无需注销",
		})
		return
	}
	checkToken := strings.SplitN(tokenHeader, " ", 2)
	if len(checkToken) != 2 && checkToken[0] != "Bearer" {
		c.JSON(http.StatusOK, gin.H{
			"stata":   "FALSE",
			"message": "token格式错误,注销失败",
		})
		c.Abort()
		return
	}
	key, err := middleware.CheckToken(checkToken[1])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"stata":   "FALSE",
			"message": "注销失败：" + err.Error(),
		})
		c.Abort()
		return
	}
	Rdb := db.NewClient()
	v, dberr := Rdb.HDEL("users", key.Num)
	if dberr != nil && v != 1 {
		c.JSON(http.StatusOK, gin.H{
			"status":  "False",
			"message": dberr.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "感谢使用！",
	})
	return

}
