package entity

import "time"

type PaymentStatus struct {
	ID            uint   `gorm:"primaryKey"`
	TransactionID uint   `gorm:""`
	Status        string `gorm:"type:enum('Pending', 'Paid', 'Failed');default:'Pending'"`
	UpdatedAt     time.Time
	Transaction   Transaction `gorm:"foreignKey:TransactionID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
