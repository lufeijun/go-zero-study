// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"demo/internal/config"
	"demo/model"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 1. 创建 sqlx.SqlConn
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	db, _ := conn.RawDB()

	db.SetMaxOpenConns(c.Mysql.MaxOpenConns)
	db.SetMaxIdleConns(c.Mysql.MaxIdleConns)

	db.SetConnMaxIdleTime(time.Duration(c.Mysql.ConnMaxIdleTime) * time.Second)
	db.SetConnMaxLifetime(time.Duration(c.Mysql.ConnMaxLifetime) * time.Second)

	fmt.Println("====NewServiceContext======")

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUsersModel(conn),
	}
}
