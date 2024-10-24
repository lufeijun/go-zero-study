package logic

import (
	"context"

	"demo/common/cryptx"
	"demo/user/model"
	"demo/user/rpc/internal/svc"
	"demo/user/rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登录逻辑
func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {

	res, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)

	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "用户不存在")
		}

		return nil, status.Error(500, err.Error())
	}

	// 判断密码
	password := cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password)

	if password != res.Password {
		return nil, status.Error(100, "密码错误")
	}

	return &user.LoginResponse{
		Id:     int64(res.Id),
		Name:   res.Name,
		Gender: int64(res.Gender),
		Mobile: res.Mobile,
	}, nil
}
