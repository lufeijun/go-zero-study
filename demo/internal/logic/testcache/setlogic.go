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
	"github.com/zeromicro/go-zero/core/fx"
	"github.com/zeromicro/go-zero/core/logx"
)

type SetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetLogic {
	return &SetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetLogic) Set() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = tool.ResponseInit()

	cache := l.ctx.Value("test_cache").(*collection.Cache)

	cache.Set("name", "吉鹏")
	cache.Set("age", 10)

	fmt.Println(cache)

	streamline()
	return
}

func streamline() {

	iter := []any{1, 2, 3, 4, 5}

	fx.Just(iter...).
		// Distinct(func(item any) any {
		// 	return item.(int) % 2
		// }).
		ForEach(func(item any) {
			fmt.Println(item)
		})

	fmt.Println("======")
}
