package logic

import (
	"context"

	"demo/study/internal/svc"
	"demo/study/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StudyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StudyLogic {
	return &StudyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StudyLogic) Study() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return &types.Response{
		Message: "Hello go-zero",
	}, nil
}
