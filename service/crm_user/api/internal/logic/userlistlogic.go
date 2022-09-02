package logic

import (
	"context"

	"demo/service/crm_user/api/internal/svc"
	"demo/service/crm_user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return &types.Response{}, nil
}
