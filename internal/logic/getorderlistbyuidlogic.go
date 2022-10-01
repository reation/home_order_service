package logic

import (
	"context"

	"github.com/reation/home_order_service/internal/svc"
	"github.com/reation/home_order_service/proto/types/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderListByUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderListByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderListByUidLogic {
	return &GetOrderListByUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderListByUidLogic) GetOrderListByUid(in *proto.UserId) (*proto.OrderListResponse, error) {
	// todo: add your logic here and delete this line
	var orderInfo proto.OrderInfo
	result := make(map[int64]*proto.OrderInfo)
	orderInfo.Id = 7
	orderInfo.Oid = 73811
	orderInfo.Uid = in.Uid
	orderInfo.Gid = 663098
	orderInfo.Price = 574.39
	result[0] = &orderInfo

	orderInfo.Id = 17
	orderInfo.Oid = 733441
	orderInfo.Uid = in.Uid
	orderInfo.Gid = 33098
	orderInfo.Price = 99.99
	result[1] = &orderInfo

	return &proto.OrderListResponse{OrderList: result}, nil
}
