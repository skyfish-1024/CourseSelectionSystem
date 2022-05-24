# 学生选课系统（**CourseSelectionSystem**）

## 一、项目简介

基于gin框架的go语言项目实践——学生选课系统

后端框架：gin框架

数据库：MySQL+redis

## 二、系统框架

![image-20220420201023349](https://user-images.githubusercontent.com/93390152/164453671-453f4271-5805-4a4f-a1ff-e7457719cf19.png)

## 三、功能

### （一）、公共功能

##### 1、登录、注销

2、查询单个学生信息

3、查询单个教师信息

### （二）、学生功能

##### 1、选课、退课、查看已选课程

2、编辑学生信息



### （三）、教师功能

##### 1、查看自己课程的学生详情

2、编辑学生信息、查询多个学生信息

3、编辑教师信息

### （四）、教务处（管理员）功能

##### 1、新增、编辑、修改课程信息（人数和时间）

2、添加（注册）学生、编辑学生信息、查询多个学生信息、删除学生

3、添加（注册）教师、编辑教师信息、查询多个教师信息、删除教师

4、添加（注册）管理员、编辑管理员信息、查询单个管理员信息、查询多个管理员信息、删除管理员

5、开放课程、关闭课程、删除某课程中的学生

## 四、表设计

#### （一）Mysql

##### 1、学生student

| 名称       | 类型        | 描述            |
| ---------- | ----------- | --------------- |
| gorm.Model |             | gorm模板        |
| Name       | varchar(20) | 学生姓名        |
| Password   | varchar(20) | 密码            |
| StuNum     | varchar(10) | 学号            |
| Grade      | int         | 年级码          |
| Class      | int         | 班级码          |
| Major      | varchar(20) | 专业            |
| Sex        | varchar(2)  | 性别            |
| IdCard     | varchar(20) | 身份证          |
| Phone      | varchar(20) | 电话            |
| Email      | varchar(20) | 邮件            |
| Role       | int         | 角色码，默认为1 |

##### 2、教师teacher

| 名称       | 类型        | 描述            |
| ---------- | ----------- | --------------- |
| gorm.Model |             | gorm模板        |
| Name       | varchar(20) | 教师姓名        |
| Password   | varchar(20) | 密码            |
| TeachNum   | varchar(10) | 教师工号        |
| Grade      | int         | 年级码          |
| Position   | varchar(20) | 职称            |
| Sex        | varchar(2)  | 性别            |
| IdCard     | varchar(20) | 身份证          |
| Phone      | varchar(20) | 电话            |
| Email      | varchar(20) | 邮箱            |
| Role       | int         | 角色码，默认为2 |

##### 3、管理员administrator

| 名称       | 类型        | 描述            |
| ---------- | ----------- | --------------- |
| gorm.Model |             | gorm模板        |
| Name       | varchar(20) | 管理员姓名      |
| Password   | varchar(20) | 密码            |
| AdminNum   | varchar(10) | 管理员工号      |
| Sex        | varchar(2)  | 性别            |
| IdCard     | varchar(20) | 身份证          |
| Phone      | varchar(20) | 电话            |
| Email      | varchar(20) | 邮箱            |
| Role       | int         | 角色码，默认为3 |

##### 4、课程course

| 名称        | 类型        | 描述                     |
| ----------- | ----------- | ------------------------ |
| gorm.Model  |             | gorm模板                 |
| CourseName  | varchar(20) | 课程名称                 |
| CourseNum   | varchar(10) | 课程码                   |
| Grade       | int         | 授课年级                 |
| Period      | int         | 学时                     |
| Score       | int         | 学分                     |
| Teacher     | varchar(20) | 任课教师                 |
| TeachNum    | varchar(10) | 教师工号                 |
| TotalStu    | int         | 学生总人数               |
| LeftStu     | int         | 学生剩余数               |
| Day         | int         | 上课星期                 |
| Time        | int         | 上课时段                 |
| Description | varchar(20) | 课程描述                 |
| State       | int         | 课程状态：1可变，2不可变 |

##### 5、选课course_choosing

| 名称       | 类别        | 描述         |
| ---------- | ----------- | ------------ |
| gorm.Model |             | gorm模板     |
| CourseName | varchar(20) | 课程名称     |
| CourseNum  | varchar(10) | 课程码       |
| Teacher    | varchar(20) | 任课教师     |
| StuName    | varchar(20) | 选课学生姓名 |
| StuNum     | varchar(10) | 选课学生学号 |
| Day        | int         | 上课星期     |
| Time       | int         | 上课时段     |

#### （二）redis

| 名称                  | 类型   | key                   | value        | 描述                 |
| --------------------- | ------ | --------------------- | ------------ | -------------------- |
| users                 | HASH   | 工号学号              | 角色码       | 记录已登录用户       |
| Stu+学号              | SET    |                       | 课程码       | 储存学生已选课程     |
| course+课程码+LeftStu | string | course+课程码+LeftStu | 剩余学生数量 | 储存课程剩余学生数量 |
| course+课程码         | SET    |                       | 学号         | 储存已选课程学生学号 |

### 五、项目结构

#### （一）、结构图

![image-20220420210440021](https://user-images.githubusercontent.com/93390152/164453814-471dd0f0-36e5-4c9d-8984-4a36d59c2b18.png)

#### （二）、目录功能

| 目录名     | 用途               | 备注 |
| ---------- | ------------------ | ---- |
| api        | 存放路由函数代码   |      |
| config     | 存放配置文件       |      |
| db         | 存放数据库相关代码 |      |
| functions  | 存放各实现功能代码 |      |
| middleware | 存放中间件相关代码 |      |
| model      | 存放模板文件       |      |
| routers    | 存放路由组         |      |
| web        | 存放前端代码       |      |

### 六、接口

测试工具：apipost

测试文档：http://124.223.70.151/z/courseselectionsystem.html

### 七、项目源码

https://github.com/zzyzzyzzyzz/CourseSelectionSystem

