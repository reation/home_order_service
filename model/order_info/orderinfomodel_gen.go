// Code generated by goctl. DO NOT EDIT!

package order_info

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	orderInfoFieldNames          = builder.RawFieldNames(&OrderInfo{})
	orderInfoRows                = strings.Join(orderInfoFieldNames, ",")
	orderInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(orderInfoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	orderInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(orderInfoFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	orderInfoModel interface {
		Insert(ctx context.Context, data *OrderInfo) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*OrderInfo, error)
		Update(ctx context.Context, data *OrderInfo) error
		Delete(ctx context.Context, id int64) error
		FindListByUID(ctx context.Context, uid int64) (*[]OrderInfo, error)
	}

	defaultOrderInfoModel struct {
		conn  sqlx.SqlConn
		table string
	}

	OrderInfo struct {
		Id            int64     `db:"id"`
		OrderNum      string    `db:"order_num"`      // 订单编号
		UId           int64     `db:"u_id"`           // 用户ID
		OrderPrice    float64   `db:"order_price"`    // 订单金额
		DiscountPrice float64   `db:"discount_price"` // 优惠金额
		RealPrice     float64   `db:"real_price"`     // 实际支付金额
		Status        int64     `db:"status"`         // 订单状态 1:待支付，2：已支付，3：取消，4：退款中，5：已退款，6：已完成
		UpdateTime    time.Time `db:"update_time"`
		CreateTime    time.Time `db:"create_time"`
	}
)

func newOrderInfoModel(conn sqlx.SqlConn) *defaultOrderInfoModel {
	return &defaultOrderInfoModel{
		conn:  conn,
		table: "`order_info`",
	}
}

func (m *defaultOrderInfoModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultOrderInfoModel) FindOne(ctx context.Context, id int64) (*OrderInfo, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", orderInfoRows, m.table)
	var resp OrderInfo
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderInfoModel) FindListByUID(ctx context.Context, id int64) (*[]OrderInfo, error) {
	query := fmt.Sprintf("select %s from %s where `u_id` = ? ", orderInfoRows, m.table)
	var resp []OrderInfo
	err := m.conn.QueryRowsCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultOrderInfoModel) Insert(ctx context.Context, data *OrderInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, orderInfoRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.OrderNum, data.UId, data.OrderPrice, data.DiscountPrice, data.RealPrice, data.Status)
	return ret, err
}

func (m *defaultOrderInfoModel) Update(ctx context.Context, data *OrderInfo) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, orderInfoRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.OrderNum, data.UId, data.OrderPrice, data.DiscountPrice, data.RealPrice, data.Status, data.Id)
	return err
}

func (m *defaultOrderInfoModel) tableName() string {
	return m.table
}