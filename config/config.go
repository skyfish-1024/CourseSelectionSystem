package config

//Mysql配置
const (
	MqRoot         = "root"                  //用户名
	MqPassword     = "密码"                    //密码
	MqAddress      = "127.0.0.1:3306"        //地址
	MqDatabaseName = "courseselectionsystem" //数据库
)

//redis配置
const (
	RedisAddr        = "localhost:6379"
	RedisIdLeTimeout = 5  // 连接等待时间
	RedisMaxIdle     = 20 // 最大空闲连接数
	RedisMaxActive   = 8  // 最大活跃连接数
	RedisPassword    = "密码"
)

//邮箱配置
const (
	EVC       = "授权码"  // EVC 邮箱授权码
	FromEmail = "qq邮箱" // FromEmail 发送方qq邮箱

)
