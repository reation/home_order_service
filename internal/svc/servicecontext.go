package svc

import (
	"github.com/reation/home_order_service/internal/config"
	"github.com/reation/home_order_service/model/order_info"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config  config.Config
	OrderDB order_info.OrderInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	OrderDB := sqlx.NewMysql(c.Mysql.OrderData.DataSourceName)
	return &ServiceContext{
		Config:  c,
		OrderDB: order_info.NewOrderInfoModel(OrderDB),
	}
}
