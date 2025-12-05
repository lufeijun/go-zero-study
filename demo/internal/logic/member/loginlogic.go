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

	var useraa Useraa
	useraa.Name = *user.Name
	useraa.Password = *user.Password
	useraa.Phone = *user.Phone

	logx.Infov(useraa)
	logx.Info(useraa)

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

type Useraa struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

func (u Useraa) MaskSensitive() any {
	return Useraa{
		Name:     u.Name,
		Password: "******",           // 密码脱敏
		Phone:    maskPhone(u.Phone), // 手机号脱敏
	}
}

// 手机号脱敏函数
func maskPhone(phone string) string {
	if len(phone) < 7 {
		return phone
	}
	return phone[:3] + "****" + phone[len(phone)-3:]
}
