package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Weather struct {
	gorm.Model
	Location    string
	Name        string
	Temperature float64
}

type WeatherResponse struct {
	ID          uint      `json:"id"`
	Location    string    `json:"location"`
	Name        string    `json:"name"`
	Temperature float64   `json:"temperature"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deteledAt,omitempty"`
}
