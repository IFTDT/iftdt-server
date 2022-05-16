package entity

import (
	"time"

	"github.com/iftdt/server/dto"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func init() {
	dto.DB.AutoMigrate(&User{})
}
