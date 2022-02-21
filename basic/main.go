package main

import "fmt"

func mapExample() {
	balance := map[string]float64{
		"623748280": 100000,
		"623343280": 200000,
	}
	fmt.Println(balance)
	fmt.Println(balance["623748280"])

}
func main() {
	mapExample()
}
