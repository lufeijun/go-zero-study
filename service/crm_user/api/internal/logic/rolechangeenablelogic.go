package logic

import (
	"context"

	"demo/service/crm_user/api/internal/svc"
	"demo/service/crm_user/api/internal/types"
	"demo/service/crm_user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleChangeEnableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleChangeEnableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleChangeEnableLogic {
	return &RoleChangeEnableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleChangeEnableLogic) RoleChangeEnable(req *types.Id) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	role, err := l.svcCtx.RoleModel.FindOne(l.ctx, req.Id)

	if err != nil {
		if err == model.ErrNotFound {
			return &types.Response{
				Status:  1,
				Message: "id 参数有误",
			}, nil
		}
		return &types.Response{
			Status:  1,
			Message: err.Error(),
		}, nil
	}

	err = l.svcCtx.RoleModel.UpdateEnable(l.ctx, req.Id, (role.IsEnable+1)%2)

	if err != nil {
		return &types.Response{
			Status:  1,
			Message: err.Error(),
		}, nil
	}

	return &types.Response{
		Status:  0,
		Message: "更新成功",
	}, nil
}
