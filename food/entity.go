package food

import "time"

type Food struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Rating      int
	CreatedAt   time.Time
	UpdateAt    time.Time
}
