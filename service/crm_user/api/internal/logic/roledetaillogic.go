package logic

import (
	"context"

	"demo/service/crm_user/api/internal/svc"
	"demo/service/crm_user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDetailLogic {
	return &RoleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDetailLogic) RoleDetail(req *types.Id) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	data, err := l.svcCtx.RoleModel.Detail(l.ctx, req.Id)

	if err != nil {
		return &types.Response{
			Status:  1,
			Message: "未找到对应记录",
		}, nil
	}

	return &types.Response{
		Data: data,
	}, nil
}
