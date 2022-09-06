package test09

import (
	"context"
	"time"

	"demo/service/test/api/internal/svc"
	"demo/service/test/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SleepLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSleepLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SleepLogic {
	return &SleepLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SleepLogic) Sleep() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	start := time.Now().Unix()

	time.Sleep(2 * time.Second)

	end := time.Now().Unix()

	return &types.Response{
		Message: "sleep",
		Data:    end - start,
	}, nil
}
