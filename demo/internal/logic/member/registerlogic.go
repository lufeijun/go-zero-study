// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package member

import (
	"context"
	"errors"
	"fmt"

	"demo/internal/svc"
	"demo/internal/types"
	"demo/model"
	"demo/tool"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

	var user model.User
	err = l.svcCtx.Db.Where("phone = ?", req.Phone).First(&user).Error

	fmt.Println(err)

	if err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		err = tool.ResponseErrorDefaultCode("手机号已存在")
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	stringPassword := string(hashedPassword)

	// 创建用户

	newobj := model.User{
		Name:     &req.Name,
		Password: &stringPassword,
		Phone:    &req.Phone,
	}
	result := l.svcCtx.Db.Create(&newobj) // 通过数据的指针来创建

	if result.Error != nil {
		return nil, result.Error
	}

	resp.Data = newobj
	return
}
