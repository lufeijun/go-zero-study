package logic

import (
	"context"

	"demo/product/model"
	"demo/product/rpc/internal/svc"
	"demo/product/rpc/product"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/status"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *product.RemoveRequest) (*product.RemoveResponse, error) {
	// todo: add your logic here and delete this line

	//	查询产品是否存在
	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == model.ErrNotFound {
			return nil, status.Error(100, "产品不存在")
		}
		return nil, status.Error(500, err.Error())
	}

	err = l.svcCtx.ProductModel.Delete(l.ctx, res.Id)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &product.RemoveResponse{}, nil
}
