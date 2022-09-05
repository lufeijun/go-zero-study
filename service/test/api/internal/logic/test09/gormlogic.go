package test09

import (
	"context"

	"demo/service/test/api/internal/svc"
	"demo/service/test/api/internal/types"
	"demo/service/test/models"

	"github.com/zeromicro/go-zero/core/logx"
)

type GormLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGormLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GormLogic {
	return &GormLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GormLogic) Gorm() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	var user models.User

	l.svcCtx.MysqlEngin.First(&user, 1)

	return &types.Response{
		Status:  0,
		Message: "success",
		Data:    user,
	}, nil
}
