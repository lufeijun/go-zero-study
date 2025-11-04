// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package member

import (
	"context"
	"errors"
	"time"

	"demo/internal/svc"
	"demo/internal/types"
	"demo/tool"

	"github.com/golang-jwt/jwt/v4"
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
	user, err := l.svcCtx.UserModel.FindOneByPhone(l.ctx, req.Phone)
	if err != nil {
		return nil, errors.New("手机号或密码错误")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("手机号或密码错误")
	}

	// 生成 JWT token
	token, err := l.generateToken(user.Id, user.Name)
	if err != nil {
		return nil, err
	}
	resp.Data = map[string]string{
		"token": token,
	}

	return
}

func (l *LoginLogic) generateToken(userId uint64, username string) (string, error) {
	now := time.Now().Unix()
	expiresAt := now + l.svcCtx.Config.JwtAuth.AccessExpire

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   userId,
		"username": username,
		"exp":      expiresAt,
		"iat":      now,
	})

	return token.SignedString([]byte(l.svcCtx.Config.JwtAuth.AccessSecret))
}
