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

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	ctx := context.Background()

	// 判断是否注册过
	_, err := l.svcCtx.UserModel.FindOneByMobile(ctx, in.Mobile)
	if err == nil {
		return nil, status.Error(100, "该用户已经存在")
	}

	if err == model.ErrNotFound {
		newUser := model.User{
			Name:     in.Name,
			Gender:   uint64(in.Gender),
			Mobile:   in.Mobile,
			Password: cryptx.PasswordEncrypt(l.svcCtx.Config.Salt, in.Password),
		}

		res, err := l.svcCtx.UserModel.Insert(ctx, &newUser)

		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		Id, err := res.LastInsertId()
		if err != nil {
			return nil, status.Error(500, err.Error())
		}

		return &user.RegisterResponse{
			Id:     Id,
			Name:   newUser.Name,
			Gender: int64(newUser.Gender),
			Mobile: newUser.Mobile,
		}, nil

	}
	return nil, status.Error(500, err.Error())
}
