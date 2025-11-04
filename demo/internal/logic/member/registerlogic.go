// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package member

import (
	"context"
	"fmt"
	"time"

	"demo/internal/svc"
	"demo/internal/types"
	"demo/model"
	"demo/tool"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = tool.ResponseInit()

	_, err = l.svcCtx.UserModel.FindOneByPhone(l.ctx, req.Phone)
	if err == nil {
		err = tool.ResponseErrorDefaultCode("手机号已存在")
		return
	}

	if err != sqlx.ErrNotFound {
		return nil, tool.ResponseErrorDefaultCode(err.Error())
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := &model.Users{
		Name:      req.Name,
		Password:  string(hashedPassword),
		Phone:     req.Phone,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	result, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()

	user.Id = uint64(id)

	fmt.Println(id)

	resp.Data = user

	return
}
