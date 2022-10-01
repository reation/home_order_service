// Code generated by goctl. DO NOT EDIT!
// Source: order.proto

package server

import (
	"context"

	"github.com/reation/home_order_service/internal/logic"
	"github.com/reation/home_order_service/internal/svc"
	"github.com/reation/home_order_service/proto/types/proto"
)

type OrderServer struct {
	svcCtx *svc.ServiceContext
	proto.UnimplementedOrderServer
}

func NewOrderServer(svcCtx *svc.ServiceContext) *OrderServer {
	return &OrderServer{
		svcCtx: svcCtx,
	}
}

func (s *OrderServer) GetOrderInfoByID(ctx context.Context, in *proto.Id) (*proto.OrderInfo, error) {
	l := logic.NewGetOrderInfoByIDLogic(ctx, s.svcCtx)
	return l.GetOrderInfoByID(in)
}

func (s *OrderServer) GetOrderInfoByOid(ctx context.Context, in *proto.OrderId) (*proto.OrderInfo, error) {
	l := logic.NewGetOrderInfoByOidLogic(ctx, s.svcCtx)
	return l.GetOrderInfoByOid(in)
}

func (s *OrderServer) GetOrderListByUid(ctx context.Context, in *proto.UserId) (*proto.OrderListResponse, error) {
	l := logic.NewGetOrderListByUidLogic(ctx, s.svcCtx)
	return l.GetOrderListByUid(in)
}