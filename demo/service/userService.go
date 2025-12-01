package service

import (
	"context"
	"demo/internal/svc"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	prx string = "jwt_"
)

type UserService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type JwtMsg struct {
	UserId    uint64 `json:"jwt_user_id"`
	UserName  string `json:"jwt_user_name"`
	ExpiresAt int64  `json:"jwt_expires_at"`
	CreateAt  uint64 `json:"jwt_create_at"`
	jwt.RegisteredClaims
}

func NewUserService(ctx context.Context, svcCtx *svc.ServiceContext) *UserService {
	return &UserService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (u *UserService) GenerateTokenMap(message map[string]any) (string, error) {
	now := time.Now().Unix()

	user_id, ok := message["user_id"].(uint64)
	if !ok {
		return "", errors.New("user_id 解析有误")
	}

	userName, ok := message["user_name"].(string)
	if !ok {
		userName = ""
	}

	expiresAt := now + u.svcCtx.Config.JwtAuth.AccessExpire
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jwt_user_id":    user_id,
		"jwt_user_name":  userName,
		"jwt_expires_at": expiresAt,
		"jwt_create_at":  now,
	})

	return token.SignedString([]byte(u.svcCtx.Config.JwtAuth.AccessSecret))
}

func (u *UserService) GenerateToken(message JwtMsg) (string, error) {
	now := time.Now().Unix()

	expiresAt := now + u.svcCtx.Config.JwtAuth.AccessExpire
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"jwt_user_id":    message.UserId,
		"jwt_user_name":  message.UserName,
		"jwt_expires_at": expiresAt,
		"jwt_create_at":  now,
	})

	return token.SignedString([]byte(u.svcCtx.Config.JwtAuth.AccessSecret))
}

func (u *UserService) GetFiledByToken(key string) any {
	return u.ctx.Value(prx + key)
}

func (u *UserService) ParseToken(tokenString string) (result *JwtMsg, err error) {

	token, err := jwt.ParseWithClaims(tokenString, &JwtMsg{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.svcCtx.Config.JwtAuth.AccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JwtMsg); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
