package logic

import (
	"context"
	"time"

	"demo/service/crm_user/api/internal/svc"
	"demo/service/crm_user/api/internal/types"
	"demo/service/crm_user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 创建时间没有插入成功
func (l *RoleAddLogic) RoleAdd(req *types.RoleAdd) (resp *types.Response, err error) {

	// 判断名称是否重复
	ishave, _ := l.svcCtx.RoleModel.IshaveByName(l.ctx, req.Name, req.Level)
	if ishave {
		return &types.Response{
			Status:  1,
			Message: "角色名称重复",
		}, nil
	}

	// todo: add your logic here and delete this line
	role := model.CrmRoles{
		Name:     req.Name,
		Level:    req.Level,
		ParentId: req.ParentId,
		IsEnable: req.IsEnable,
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdateAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	res, err := l.svcCtx.RoleModel.Insert(l.ctx, &role)
	if err != nil {
		return &types.Response{
			Status:  1,
			Message: err.Error(),
		}, nil
	}

	role.Id, _ = res.LastInsertId()

	return &types.Response{
		Message: "success",
		Data:    role,
	}, nil
}
