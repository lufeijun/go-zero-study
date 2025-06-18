package logic

import (
	"context"

	"demo/demo/internal/svc"
	"demo/demo/internal/tools"
	"demo/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoLogic) Demo() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = tools.ResponseInit()
	resp.Message = "Hello go-zero"

	// err = errors.New("测试")

	l.Logger.Info("====测试")

	return
}
