package logic

import (
	"context"
	"fmt"

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

	orderLists, err := l.svcCtx.OrderDB.FindListByUID(l.ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	orderList := *orderLists
	fmt.Println(orderList)
	result := make(map[int64]*proto.OrderInfo)
	for k, v := range orderList {
		var orderInfo proto.OrderInfo
		orderInfo.Id = v.Id
		orderInfo.Oid = v.Id
		orderInfo.Uid = v.UId
		orderInfo.Gid = v.UId
		orderInfo.Price = float32(v.RealPrice)
	}
	fmt.Println(result)
	return &proto.OrderListResponse{OrderList: result}, nil
}
