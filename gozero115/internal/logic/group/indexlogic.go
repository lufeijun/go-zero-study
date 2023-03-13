package group

import (
	"context"
	"fmt"
	"time"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"
	"demo/gozero115/tool/funcs"

	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/logx"
)

type IndexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIndexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IndexLogic {
	return &IndexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IndexLogic) Index() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	resp = funcs.ResponseInit()

	timeWheel, err := collection.NewTimingWheel(time.Second, 300, func(key, value any) {
		fmt.Println(time.Now().UnixMilli())
		fmt.Println(key, "===ok===", value)
	})

	if err != nil {
		fmt.Println("error")
	}

	timeWheel.SetTimer("aa", "bbb", 3*time.Second)
	fmt.Println(time.Now().UnixMilli())

	return
}
