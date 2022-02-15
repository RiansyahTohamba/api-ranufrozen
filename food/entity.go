package food

import (
	"time"
)

type Food struct {
	ID            int
	Name          string
	PhotoPath     string
	Rating        int
	Price         float64
	Stock         int
	IsSuperSeller int8
	Category      int
	QuantitySold  int

	Description string
	Discount    float32

	CreatedAt time.Time
	UpdateAt  time.Time
}
