package order

import (
	"fmt"
	"time"
)

type Order struct {
	Id   int
	Date time.Time
	Sum  float64
}

func (order Order) display() {
	fmt.Println("sum: %s", order.Sum)
}

func main() {
	order1 := Order{
		Id:  1,
		Sum: 1000.0,
	}
	order1.display()
}
