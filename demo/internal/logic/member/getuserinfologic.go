// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package member

import (
	"context"
	"fmt"

	"demo/internal/svc"
	"demo/internal/types"
	"demo/model"
	"demo/service"
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

	userService := service.NewUserService(l.ctx, l.svcCtx)
	msg, _ := userService.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqd3RfY3JlYXRlX2F0IjoxNzY0NTkzODA5LCJqd3RfZXhwaXJlc19hdCI6MTc2NDY4MDIwOSwiand0X3VzZXJfaWQiOjEsImp3dF91c2VyX25hbWUiOiLot6_po57lkJsifQ.Dss4YZ9uxjvwJKPVrEfdSetx7IPT3dC2cIaOWpe1MKw")
	fmt.Println(msg)
	fmt.Println(msg.UserName)

	fmt.Println(l.ctx)

	return
}
