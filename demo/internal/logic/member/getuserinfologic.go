// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package member

import (
	"context"

	"demo/internal/svc"
	"demo/internal/types"
	"demo/model"
	"demo/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfo) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = tool.ResponseInit()

	var user model.User

	err = l.svcCtx.Db.First(&user, req.ID).Error

	if err != nil {
		return nil, err
	}
	resp.Data = user
	return
}
