package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type WithdrawOrder struct {
	Id           int64           `json:"id" gorm:"primaryKey;autoIncrement;comment:提现订单ID"`
	UserId       int64           `json:"userId" gorm:"column:user_id;type:bigint;comment:用户ID"`
	Amount       decimal.Decimal `json:"amount" gorm:"column:amount;type:decimal(30,2);comment:提现金额"`
	Address      string          `json:"address" gorm:"column:address;type:varchar(255);comment:提现地址"`
	Free         decimal.Decimal `json:"free" gorm:"column:free;type:decimal(30,2);comment:提现手术费固定0.03"`
	Status       int64           `json:"status" gorm:"column:status;type:tinyint;comment:状态 1待审核 2通过 3拒绝 4打款完成"`
	RejectReason string          `json:"rejectReason" gorm:"column:reject_reason;type:varchar(255);comment:拒绝原因"`
	CreatedAt    *time.Time      `json:"createdAt" gorm:"column:created_at;type:datetime;comment:申请时间"`
	ReviewedAt   *time.Time      `json:"reviewedAt" gorm:"column:reviewed_at;type:datetime;comment:审核时间"`
}

func (WithdrawOrder) TableName() string {
	return "app_withdraw_order"
}
