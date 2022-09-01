package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CrmRolesModel = (*customCrmRolesModel)(nil)

type (
	// CrmRolesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCrmRolesModel.
	CrmRolesModel interface {
		crmRolesModel
	}

	customCrmRolesModel struct {
		*defaultCrmRolesModel
	}
)

// NewCrmRolesModel returns a model for the database table.
func NewCrmRolesModel(conn sqlx.SqlConn) CrmRolesModel {
	return &customCrmRolesModel{
		defaultCrmRolesModel: newCrmRolesModel(conn),
	}
}
