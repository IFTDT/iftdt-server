package entity

import (
	"time"

	"github.com/iftdt/server/dto"
)

type Device struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeviceToken string    `json:"device_token" gorm:"unique"`

	UserId uint `json:"user_id"`
	User   User `json:"-"`
}

func init() {
	dto.DB.AutoMigrate(&Device{})
}
