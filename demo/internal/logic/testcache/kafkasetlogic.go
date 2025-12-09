// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package testcache

import (
	"context"
	"encoding/json"
	"fmt"

	"demo/internal/svc"
	"demo/internal/types"
	"demo/model"
	"demo/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaSetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKafkaSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *KafkaSetLogic {
	return &KafkaSetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *KafkaSetLogic) KafkaSet() (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	resp = tool.ResponseInit()
	resp.Data = "kafka set test success"

	name := "zhangSan"
	phone := "13800138000"
	users := model.User{
		Name:  &name,
		Phone: &phone,
	}

	// 结构体 → JSON ([]byte)
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if err := l.svcCtx.KqPusherClient.Push(l.ctx, string(jsonBytes)); err != nil {
		
		logx.Errorf("KqPusherClient Push Error , err :%v", err)
	}

	return
}
