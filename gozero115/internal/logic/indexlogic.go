package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"demo/gozero115/internal/svc"
	"demo/gozero115/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
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
	// var a interface{}
	start := time.Now().UnixMicro()

	time.Sleep(time.Second)

	// mrFinish()

	return &types.Response{
		Data: time.Now().UnixMicro() - start,
	}, nil
}

func mrFinish() {
	err := mr.Finish(func() (err error) {
		time.Sleep(time.Millisecond * 3)
		return
	}, func() error {
		time.Sleep(time.Millisecond * 2)
		return errors.New("报错了")
	}, func() error {
		time.Sleep(time.Millisecond * 1)
		return nil
	})

	if err != nil {
		fmt.Println("ok")
	}
}
