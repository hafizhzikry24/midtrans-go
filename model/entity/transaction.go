package entity

import "time"

type Transaction struct {
	ID          uint   `gorm:"primaryKey"`
	OrderID     string `gorm:"unique"`
	UserID      int    `gorm:"foreignKey:UserID"`
	ItemID      string
	ItemName    string
	Amount      int64
	Token       string
	RedirectUrl string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	User        User `gorm:"foreignKey:UserID"`
}
