package db

import (
	"CourseSelectionSystem/config"
	"errors"
	"github.com/garyburd/redigo/redis"
	"time"
)

//OptionPool 连接池
type OptionPool struct {
	addr        string
	idLeTimeout int
	maxIdle     int
	maxActive   int
	password    string
}

type PoolExt interface {
	apply(*OptionPool)
}

type tempFunc func(pool *OptionPool)

type funcPoolExt struct {
	f tempFunc
}

func (f *funcPoolExt) apply(p *OptionPool) {
	f.f(p)
}
func NewFuncPoolExt(f tempFunc) *funcPoolExt {
	return &funcPoolExt{f: f}
}

type Client struct {
	Option OptionPool
	pool   *redis.Pool
}

// DefaultOption 默认配置
var DefaultOption = OptionPool{
	addr:        config.RedisAddr,
	idLeTimeout: config.RedisIdLeTimeout,
	maxIdle:     config.RedisMaxIdle,
	maxActive:   config.RedisMaxActive,
	password:    config.RedisPassword,
}

//NewClient 新建连接，返回连接对象
func NewClient(op ...PoolExt) *Client {
	c := &Client{Option: DefaultOption}
	for _, p := range op {
		p.apply(&c.Option)
	}
	c.setRedisPool()
	return c
}

//setRedisPool 设置连接池
func (Rdb *Client) setRedisPool() {
	Rdb.pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", Rdb.Option.addr)
			if err != nil {
				return nil, err
			}
			_, err = conn.Do("AUTH", config.RedisPassword)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
		MaxIdle:     Rdb.Option.maxIdle,                                  // 最大空闲连接数
		MaxActive:   Rdb.Option.maxActive,                                // 最大活跃连接数
		IdleTimeout: time.Second * time.Duration(Rdb.Option.idLeTimeout), // 连接等待时间
	}
}

//OpenCourse 管理员开放可选课程
func (Rdb *Client) OpenCourse(CourseNum string, LeftStu int) (interface{}, error) {
	c := Rdb.pool.Get()
	defer c.Close()
	v, err := c.Do("SET", "course"+CourseNum+"LeftStu", LeftStu)
	if err != nil {
		return nil, err
	}
	return v, nil
}

//GetStuCourse 查询学生已选课程
func (Rdb *Client) GetStuCourse(StuNum string) ([]string, error) {
	c := Rdb.pool.Get()
	defer c.Close()
	//读取redis中课程包含的学生数据
	students, err := redis.Strings(c.Do("SMEMBERS", "Stu"+StuNum))
	if err != nil {
		return nil, err
	}
	return students, nil
}

//CloseCourse 管理员关闭可选课程
func (Rdb *Client) CloseCourse(CourseNum string) ([]string, error) {
	c := Rdb.pool.Get()
	defer c.Close()
	//读取redis中课程包含的学生数据
	students, err := redis.Strings(c.Do("SMEMBERS", "course"+CourseNum))
	if err != nil {
		return nil, err
	}
	return students, nil
}

//CourseAddStu 学生选择课程，将已选学生写入课程集合
func (Rdb *Client) CourseAddStu(CourseNum string, StuNum string) error {
	c := Rdb.pool.Get()
	defer c.Close()
	//判断学生是否在集合中，若在，则取消，避免重复加入
	v, err := c.Do("SISMEMBER", "course"+CourseNum, StuNum)
	if err != nil {
		return err
	}
	if v.(int64) == 1 {
		return errors.New("重复选择")
	}
	//在课程中写入学生
	v, err = c.Do("SADD", "course"+CourseNum, StuNum)
	if err != nil {
		return err
	}
	//在学生中写入课程
	v, err = c.Do("SADD", "Stu"+StuNum, CourseNum)
	if err != nil {
		return err
	}
	//学生剩余数减一
	v, err = c.Do("DECRBY", "course"+CourseNum+"LeftStu", 1)
	if err != nil {
		return err
	}

	return nil
}

//CourseDelStu 学生退出课程，将已退学生移除课程集合
func (Rdb *Client) CourseDelStu(CourseNum string, StuNum string) (interface{}, error) {
	c := Rdb.pool.Get()
	defer c.Close()
	//判断学生是否在集合中，若不在，则返回错误
	v, err := c.Do("SISMEMBER", "course"+CourseNum, StuNum)
	if err != nil {
		return nil, err
	} else if v.(int64) == 0 {
		return v, errors.New("还未选择该课程选择")
	}
	//移除学生
	v, err = c.Do("SREM", "course"+CourseNum, StuNum)
	if err != nil {
		return nil, err
	}
	//移除学生的课程
	v, err = c.Do("SREM", "Stu"+StuNum, CourseNum)
	if err != nil {
		return nil, err
	}
	//课程剩余数加一
	v, err = c.Do("INCRBY", CourseNum+"LeftStu", 1)
	if err != nil {
		return nil, err
	}
	return v, nil
}

//VCset 设置用户验证码
func (Rdb *Client) VCset(key string, value string) error {
	c := Rdb.pool.Get()
	defer c.Close()
	_, err := c.Do("SET", "VC"+key, value, "EX", 600)
	if err != nil {
		return err
	}
	return nil
}

//VCget 获取用户验证码
func (Rdb *Client) VCget(key string) (interface{}, error) {
	c := Rdb.pool.Get()
	defer c.Close()
	v, err := redis.String(c.Do("GET", "VC"+key))
	if err != nil {
		return nil, err
	}
	return v, nil
}

//HSET 设置哈希值
func (Rdb *Client) HSET(HashName string, key string, Value interface{}) (interface{}, error) {
	c := Rdb.pool.Get()
	defer c.Close()
	v, err := c.Do("HSET", HashName, key, Value)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// HGET 获取哈希值
func (Rdb *Client) HGET(HashName string, key string) (interface{}, error) {
	c := Rdb.pool.Get()
	defer c.Close()
	v, err := redis.String(c.Do("HGET", HashName, key))
	if err != nil {
		return nil, err
	}
	return v, nil
}

//HDEL 删除哈希值
func (Rdb *Client) HDEL(HashName string, key string) (interface{}, error) {
	c := Rdb.pool.Get()
	defer c.Close()
	v, err := c.Do("HDEL", HashName, key)
	if err != nil {
		return nil, err
	}
	return v, nil
}
