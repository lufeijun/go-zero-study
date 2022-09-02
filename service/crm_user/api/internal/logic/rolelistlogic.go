package logic

import (
	"context"

	"demo/service/crm_user/api/internal/svc"
	"demo/service/crm_user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.RoleSearch) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	result, err := l.svcCtx.RoleModel.List(l.ctx, req.Name, req.Level, req.IsEnable, req.ParentId, req.Page)

	if err != nil {
		return &types.Response{
			Status:  1,
			Message: err.Error(),
		}, nil
	}

	return &types.Response{
		Message: "success",
		Data:    result,
	}, nil
}
