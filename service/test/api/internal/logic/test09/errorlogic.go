package test09

import (
	"context"
	"errors"

	"demo/service/test/api/internal/svc"
	"demo/service/test/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ErrorLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewErrorLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ErrorLogic {
	return &ErrorLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ErrorLogic) Error(req *types.ErrorReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	i := 1 / req.A

	return nil, errors.New("报错")

	return &types.Response{
		Message: "error",
		Data:    i,
	}, nil
}
