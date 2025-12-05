// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"demo/internal/config"
	"demo/internal/middleware"
	"time"

	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config config.Config

	Db *gorm.DB

	// 中间件
	TestCacheMiddleware rest.Middleware

	SingleFlight syncx.SingleFlight
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 初始化 GORM
	db, err := gorm.Open(mysql.Open(c.Mysql.DataSource), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// 获取底层的 sql.DB 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		panic("failed to get sql.DB: " + err.Error())
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(c.Mysql.MaxIdleConns)                                    // 最大空闲连接数
	sqlDB.SetMaxOpenConns(c.Mysql.MaxOpenConns)                                    // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Duration(c.Mysql.ConnMaxLifetime) * time.Second) // 连接最大存活时间
	sqlDB.SetConnMaxIdleTime(time.Duration(c.Mysql.ConnMaxIdleTime) * time.Second)

	return &ServiceContext{
		Config: c,
		Db:     db,

		TestCacheMiddleware: middleware.NewTestCacheMiddleware().Handle,

		SingleFlight: syncx.NewSingleFlight(),
	}
}
