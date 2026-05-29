package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// InitDB 初始化数据库连接
func InitDB() error {
	if AppConfig == nil {
		return fmt.Errorf("config not loaded")
	}

	var err error

	// 根据驱动类型选择数据库
	switch AppConfig.Database.Driver {
	case "sqlite3":
		log.Println("Using SQLite database:", AppConfig.Database.DSN)
		DB, err = gorm.Open(sqlite.Open(AppConfig.Database.DSN), &gorm.Config{
			Logger: getLogger(),
		})
	case "mysql":
		dsn := AppConfig.Database.GetDSN()
		log.Println("Using MySQL database:", dsn)
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: getLogger(),
		})
	default:
		return fmt.Errorf("unsupported database driver: %s", AppConfig.Database.Driver)
	}

	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}

	// 获取通用数据库对象 sql.DB 以设置连接池
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %v", err)
	}

	// 设置连接池（仅对MySQL有效，SQLite会忽略）
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生命周期

	log.Println("Database connected successfully")
	return nil
}

// getLogger 获取GORM日志配置
func getLogger() logger.Interface {
	var logLevel logger.LogLevel

	if AppConfig.Server.IsDebug() {
		logLevel = logger.Info
	} else {
		logLevel = logger.Warn
	}

	return logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second, // 慢SQL阈值
			LogLevel:                  logLevel,    // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound
			Colorful:                  true,        // 彩色打印
		},
	)
}

// GetDB 获取数据库实例
func GetDB() *gorm.DB {
	return DB
}
