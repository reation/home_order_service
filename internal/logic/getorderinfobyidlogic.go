package logic

import (
	"context"

	"github.com/reation/home_order_service/internal/svc"
	"github.com/reation/home_order_service/proto/types/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderInfoByIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderInfoByIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderInfoByIDLogic {
	return &GetOrderInfoByIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderInfoByIDLogic) GetOrderInfoByID(in *proto.Id) (*proto.OrderInfo, error) {
	// todo: add your logic here and delete this line

	return &proto.OrderInfo{
		Id:    in.Id,
		Oid:   77858,
		Uid:   7655,
		Gid:   44325,
		Price: 59.99,
	}, nil
}
