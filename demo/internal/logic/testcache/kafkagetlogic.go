// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package testcache

import (
	"context"

	"demo/internal/svc"
	"demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaGetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKafkaGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KafkaGetLogic {
	return &KafkaGetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KafkaGetLogic) KafkaGet() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
