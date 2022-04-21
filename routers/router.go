package routers

import (
	"CourseSelectionSystem/api/v1"
	"CourseSelectionSystem/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	//登录与注销
	r.POST("/login", v1.Login)
	r.GET("/logout", v1.Logout)

	//公共路由
	//找回密码
	r.POST("/find_pwd/sendmail", v1.SendEmail)
	r.POST("/find_pwd/vc_check", v1.VCCheck)
	r.POST("/find_pwd/reset_pwd", middleware.PRTJwt(), v1.PasswordReset)
	//查询课程信息
	r.GET("/getCourse", v1.GetCourse)
	r.GET("/getCourses", v1.GetCourses)
	//查询学生信息
	r.GET("/getStudent", v1.GetStudent)
	//查询教师信息
	r.GET("/getTeacher", v1.GetTeacher)

	//学生路由组
	S := r.Group("student")
	S.Use(middleware.StuJwt())
	{
		//学生相关路由
		S.PUT("/editStudent", v1.EditeStudent)

		//课程相关路由
		S.POST("/chooseCourse", v1.CourseChoose)
		S.DELETE("/dropCourse", v1.CourseDrop)
		S.GET("/getChoseCourses", v1.GetChoseCourses)
	}

	//教师路由组
	T := r.Group("teacher")
	T.Use(middleware.TechJwt())
	{
		//学生相关路由
		T.PUT("editStudent", v1.EditeStudent)

		//教师相关路由
		T.PUT("editTeacher", v1.EditeTeacher)

		//课程相关路由
		T.GET("/getCourseChoseStu", v1.GetCourseChoseStu)
	}

	//管理员路由组
	A := r.Group("administrator")
	A.Use(middleware.AdminJwt())
	{
		//学生相关路由
		A.POST("/addStudent", v1.AddStudent)
		A.PUT("editStudent", v1.EditeStudent)
		A.DELETE("/deleteStudent", v1.DeleteStudent)
		A.GET("/getStudents", v1.GetStudents)

		//教师相关路由
		A.POST("/addTeacher", v1.AddTeacher)
		A.DELETE("/deleteTeacher", v1.DeleteTeacher)
		A.PUT("editTeacher", v1.EditeTeacher)
		A.GET("/getTeachers", v1.GetTeachers)

		//管理员相关路由
		A.POST("/addAdministrator", v1.AddAdministrator)
		A.DELETE("/deleteAdministrator", v1.DeleteAdministrator)
		A.PUT("editAdministrator", v1.EditeAdministrator)
		A.GET("/getAdministrator", v1.GetAdministrator)
		A.GET("/getAdministrators", v1.GetAdministrators)

		//课程相关路由
		A.POST("/addCourse", v1.AddCourse)
		A.DELETE("/deleteCourse", v1.DeleteCourse)
		A.PUT("editCourse", v1.EditeCourse)

		//选课相关路由
		A.POST("/openCourse", v1.OpenCourse)
		A.POST("/closeCourse", v1.CloseCourse)
		A.POST("/addCourseStu", v1.CourseAddStudent)
		A.DELETE("/dropCourseStu", v1.CourseDelStudent)
	}

	err := r.Run(":3000")
	if err != nil {
		return
	}
}
