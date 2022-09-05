package svc

import (
	"demo/service/test/api/internal/config"
	"demo/service/test/api/internal/middleware"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config          config.Config
	MysqlEngin      *gorm.DB
	Errormiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 数据库
	db, err := gorm.Open(mysql.Open(c.DataSourceName), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:          c,
		MysqlEngin:      db,
		Errormiddleware: middleware.NewErrormiddlewareMiddleware().Handle,
	}
}
