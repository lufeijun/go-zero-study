package svc

import (
	"demo/service/test/api/internal/config"
	"demo/service/test/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.UsersModel
	UserModel1 model.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {

	conn := sqlx.NewMysql(c.Mysql.DataSource)
	conn1 := sqlx.NewMysql(c.Mysql.DataSource1)

	return &ServiceContext{
		Config:     c,
		UserModel:  model.NewUsersModel(conn),
		UserModel1: model.NewUsersModel(conn1),
	}
}
