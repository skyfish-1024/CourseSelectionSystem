package config

//Mysql配置
const (
	MqRoot         = "root"                  //用户名
	MqPassword     = "13896764180zy"         //密码
	MqAddress      = "127.0.0.1:3306"        //地址
	MqDatabaseName = "courseselectionsystem" //数据库
)

//redis配置
const (
	RedisAddr        = "localhost:6379"
	RedisIdLeTimeout = 5  // 连接等待时间
	RedisMaxIdle     = 20 // 最大空闲连接数
	RedisMaxActive   = 8  // 最大活跃连接数
	RedisPassword    = "13896764180zy"
)

//邮箱配置
const (
	EVC       = "yhvwoqukrsnmchfj"  // EVC 邮箱授权码
	FromEmail = "3574819459@qq.com" // FromEmail 发送方qq邮箱

)
