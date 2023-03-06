package config

// Config 结构体用于存储配置数据
type Config struct {
	Port      int
	Database  string
	RedisHost string
	RedisPort int
	RedisPwd  string
}
