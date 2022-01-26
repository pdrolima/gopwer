package entity

import (
	"github.com/google/uuid"
)

type Order struct {
	ID              uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	OrderId         string    `json:"order_id"`
	UserId          string    `json:"user_id"`
	Amount          int       `json:"amount"`
	CashAmount      int       `json:"cash_amount"`
	Gateway         string    `json:"gateway"`
	OrderStatus     string    `json:"order_status"`
	TransactionCode string    `json:"transaction_code"`
	DeliveredAt     string    `json:"delivered_at"`
	CreatedAt       string    `json:"created_at"`
	UpdatedAt       string    `json:"updated_at"`
}
