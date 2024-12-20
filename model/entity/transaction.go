package entity

import "time"

type Transaction struct {
	ID          uint   `gorm:"primaryKey"`
	OrderID     string `gorm:"unique"`
	UserID      int
	ItemID      string
	ItemName    string
	Amount      int64
	Token       string
	RedirectUrl string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
