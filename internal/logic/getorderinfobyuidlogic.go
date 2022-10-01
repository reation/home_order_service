package logic

import (
	"context"

	"github.com/reation/home_order_service/internal/svc"
	"github.com/reation/home_order_service/proto/types/proto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderInfoByUidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderInfoByUidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderInfoByUidLogic {
	return &GetOrderInfoByUidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderInfoByUidLogic) GetOrderInfoByUid(in *proto.UserId) (*proto.OrderInfo, error) {
	// todo: add your logic here and delete this line

	return &proto.OrderInfo{}, nil
}
