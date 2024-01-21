package config

type ServerConfig struct {
	Name string `mapstructure:"name"`
	Port int    `mapstructure:"port"`
	//Mysqlinfo   MysqlConfig `mapstructure:"mysql"`
	//RedisInfo   RedisConfig `mapstructure:"redis"`
	LogsAddress string `mapstructure:"logsAddress"`
	//JWTKey      JWTConfig   `mapstructure:"jwt"`
	//MinioInfo   MinioConfig `mapstructure:"minio"`
}
