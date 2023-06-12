package logic

import (
	"context"

	"study/docker/internal/svc"
	"study/docker/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeLogic {
	return &HomeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeLogic) Home() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = &types.Response{}
	resp.Message = "success"
	resp.Status = 0
	resp.Data = "Hello World"

	return
}
