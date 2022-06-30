package food

import (
	"time"
)

type Food struct {
	Name          string  `bson:"name"`
	PhotoPath     string  `bson:"photoPath"`
	Rating        int     `bson:"rating"`
	Price         float64 `bson:"price"`
	Stock         int     `bson:"stock"`
	IsSuperSeller int8    `bson:"isSuperSeller"`
	Category      int     `bson:"category"`
	QuantitySold  int     `bson:"qtSold"`
	Description   string  `bson:"desc"`
	Discount      float32

	CreatedAt time.Time
	UpdateAt  time.Time
}
