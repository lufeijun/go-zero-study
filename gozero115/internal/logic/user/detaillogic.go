package user

import (
	"context"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = new(types.Response)

	resp.Message = "detail"

	return
}
