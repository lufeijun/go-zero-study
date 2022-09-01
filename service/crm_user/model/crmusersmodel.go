package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ CrmUsersModel = (*customCrmUsersModel)(nil)

type (
	// CrmUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCrmUsersModel.
	CrmUsersModel interface {
		crmUsersModel
	}

	customCrmUsersModel struct {
		*defaultCrmUsersModel
	}
)

// NewCrmUsersModel returns a model for the database table.
func NewCrmUsersModel(conn sqlx.SqlConn) CrmUsersModel {
	return &customCrmUsersModel{
		defaultCrmUsersModel: newCrmUsersModel(conn),
	}
}
