package tool

import (
	"context"
	"demo/internal/config"
	"demo/internal/svc"
	"demo/model"
	"encoding/json"
	"fmt"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

func Consumers(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.KqConsumerConf, NewPaymentSuccess(ctx, svcContext)),
		//.....
	}

}

type PaymentSuccess struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentSuccess(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentSuccess {
	return &PaymentSuccess{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentSuccess) Consume(ctx context.Context, key, val string) error {
	logx.Infof("PaymentSuccess key :%s , val :%s", key, val)

	var user model.User
	err := json.Unmarshal([]byte(val), &user)
	if err != nil {
		panic("解析错误:" + err.Error())
	}

	// 打印user所有字段
	fmt.Printf("User Struct: %+v\n", user)

	return nil
}
