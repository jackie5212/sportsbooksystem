package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// Config 应用配置结构
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	WeChat   WeChatConfig   `mapstructure:"wechat"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"` // debug, release, test
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Driver   string `mapstructure:"driver"` // sqlite3 or mysql
	DSN      string `mapstructure:"dsn"`    // SQLite DSN or MySQL connection string
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
	Charset  string `mapstructure:"charset"`
}

// JWTConfig JWT配置
type JWTConfig struct {
	Secret     string `mapstructure:"secret"`
	ExpireHour int    `mapstructure:"expire_hour"`
}

// WeChatConfig 微信配置
type WeChatConfig struct {
	AppID     string `mapstructure:"app_id"`
	AppSecret string `mapstructure:"app_secret"`
	MchID     string `mapstructure:"mch_id"`
	APIKey    string `mapstructure:"api_key"`
}

var AppConfig *Config

// LoadConfig 加载配置
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")

	// 设置默认值
	setDefaults()

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Warning: config file not found, using environment variables: %v\n", err)
	}

	// 从环境变量读取(优先级更高)
	viper.AutomaticEnv()

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %v", err)
	}

	AppConfig = config
	return config, nil
}

// setDefaults 设置默认配置值
func setDefaults() {
	// Server
	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.mode", "debug")

	// Database
	viper.SetDefault("database.driver", "sqlite3")
	viper.SetDefault("database.dsn", "tennis_booking.db")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "3306")
	viper.SetDefault("database.user", "root")
	viper.SetDefault("database.password", "root")
	viper.SetDefault("database.dbname", "tennis_booking")
	viper.SetDefault("database.charset", "utf8mb4")

	// JWT
	viper.SetDefault("jwt.secret", "tennis-booking-secret-key-2026")
	viper.SetDefault("jwt.expire_hour", 72)

	// WeChat
	viper.SetDefault("wechat.app_id", "")
	viper.SetDefault("wechat.app_secret", "")
	viper.SetDefault("wechat.mch_id", "")
	viper.SetDefault("wechat.api_key", "")
}

// GetDSN 获取数据库连接字符串
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DBName,
		c.Charset,
	)
}

// IsDebug 判断是否为调试模式
func (c *ServerConfig) IsDebug() bool {
	return c.Mode == "debug"
}

// GetJWTSecret 获取JWT密钥
func GetJWTSecret() []byte {
	if AppConfig != nil {
		return []byte(AppConfig.JWT.Secret)
	}
	return []byte(os.Getenv("JWT_SECRET"))
}
