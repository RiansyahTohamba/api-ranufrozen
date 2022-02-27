package main

import "fmt"

func greetBio(bioChan chan string) {
	bioChan <- "Bio"
}

func greetHani(haniChan chan string) {
	haniChan <- "Hani"
}

func mapExample() {
	balance := map[string]float64{
		"623748280": 100000,
		"623343280": 200000,
	}
	fmt.Println(balance)
	fmt.Println(balance["623748280"])

}

func main() {
	bioChannel := make(chan string)
	haniChannel := make(chan string)
	go greetBio(bioChannel)
	go greetHani(haniChannel)
	fmt.Println(<-bioChannel)
	fmt.Println(<-haniChannel)
}
