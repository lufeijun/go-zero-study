package svc

import (
	"demo/service/crm_user/api/internal/config"
	"demo/service/crm_user/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.CrmUsersModel
	RoleModel model.CrmRolesModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	return &ServiceContext{
		Config:    c,
		UserModel: model.NewCrmUsersModel(conn),
		RoleModel: model.NewCrmRolesModel(conn),
	}
}
