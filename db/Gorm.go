package db

import (
	"CourseSelectionSystem/config"
	"CourseSelectionSystem/model"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var Mdb *gorm.DB

func InitDb() {
	var err error
	Mdb, err = gorm.Open("mysql", config.MqRoot+":"+config.MqPassword+"@tcp("+config.MqAddress+")/"+config.MqDatabaseName+"?"+"charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败，请检查连接参数err:", err)
	}
	//禁用默认表名的复数形式
	Mdb.SingularTable(true)
	//自动迁移表
	Mdb.AutoMigrate(&model.Student{}, &model.Teacher{}, &model.Administrator{}, &model.Course{}, &model.CourseChoosing{})
	//最大空闲
	Mdb.DB().SetMaxIdleConns(100)
	//最大连接
	Mdb.DB().SetMaxOpenConns(100)
	//最大可复用时间
	Mdb.DB().SetConnMaxLifetime(10 * time.Second)
	//Mdb.Close()
	return
}
