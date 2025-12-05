// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package testcache

import (
	"context"
	"fmt"

	"demo/internal/svc"
	"demo/internal/types"
	"demo/tool"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLogic) Get() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = tool.ResponseInit()

	cache := l.ctx.Value("test_cache").(*collection.Cache)

	name, ok := cache.Get("name")

	if !ok {
		fmt.Println("key not found")
		return
	} else {
		namestr := name.(string)

		fmt.Println("=====")
		fmt.Println(namestr)
	}

	return
}
