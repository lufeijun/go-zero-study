package test09

import (
	"context"

	"demo/service/test/api/internal/svc"
	"demo/service/test/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Test09Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTest09Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Test09Logic {
	return &Test09Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// mysql 多个连接地址
func (l *Test09Logic) Test09() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	// user, err := l.svcCtx.UserModel.FindOne(l.ctx, 1)

	user, err := l.svcCtx.UserModel1.FindOne(l.ctx, 1)

	if err != nil {
		return &types.Response{
			Status:  1,
			Message: err.Error(),
		}, nil
	}

	return &types.Response{
		Status:  1,
		Message: "error",
		Data:    user,
	}, nil
}
