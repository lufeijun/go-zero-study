// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package member

import (
	"context"
	"errors"

	"demo/internal/svc"
	"demo/internal/types"
	"demo/model"
	"demo/service"
	"demo/tool"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = tool.ResponseInit()

	// 查找用户
	var user model.User

	err = l.svcCtx.Db.Where("phone = ?", req.Phone).First(&user).Error

	if err != nil {
		return nil, errors.New("手机号或密码错误")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("手机号或密码错误")
	}

	// 生成 JWT token
	userService := service.NewUserService(l.ctx, l.svcCtx)
	token, err := userService.GenerateToken(service.JwtMsg{
		UserId:   uint64(user.ID),
		UserName: *user.Name,
	})
	if err != nil {
		return nil, err
	}
	resp.Data = map[string]string{
		"token": token,
	}

	return
}
