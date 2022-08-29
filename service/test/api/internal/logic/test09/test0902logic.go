package test09

import (
	"context"

	"demo/service/test/api/internal/svc"
	"demo/service/test/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Test0902Logic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTest0902Logic(ctx context.Context, svcCtx *svc.ServiceContext) *Test0902Logic {
	return &Test0902Logic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 事务
func (l *Test0902Logic) Test0902() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	// err := l.svcCtx.

	return &types.Response{}, nil
}
