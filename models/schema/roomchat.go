package schema

import "time"

type Roomchat struct {
	ID             uint `gorm:"primaryKey"`
	TransactionID  uint
	Status         bool
	ExpirationTime *time.Time
	CreatedAt      time.Time
	Message        []Message `gorm:"ForeignKey:RoomchatID;references:ID"` // one to many
}
