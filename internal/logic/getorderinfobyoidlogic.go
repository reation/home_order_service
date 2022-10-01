package logic

import (
	"context"

	"github.com/reation/home_order_service/internal/svc"
	"github.com/reation/home_order_service/proto/types/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderInfoByOidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderInfoByOidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderInfoByOidLogic {
	return &GetOrderInfoByOidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderInfoByOidLogic) GetOrderInfoByOid(in *proto.OrderId) (*proto.OrderInfo, error) {
	// todo: add your logic here and delete this line

	return &proto.OrderInfo{
		Id:    321444,
		Oid:   in.Oid,
		Uid:   48832,
		Gid:   90022,
		Price: 89.99,
	}, nil
}
