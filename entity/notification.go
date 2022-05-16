package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/iftdt/server/dto"
)

type MyMap map[string]interface{}

func (c MyMap) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *MyMap) Scan(src interface{}) error {
	return json.Unmarshal(src.([]byte), c)
}

type Notification struct {
	ID        uint      `json:"id"`
	Payload   MyMap     `json:"payload" gorm:"type:json"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserID uint `json:"user_id"`
	User   User `json:"user"`
}

func init() {
	dto.DB.AutoMigrate(&Notification{})
}
