package basic

import "fmt"
/*
keyword untuk concurency
1. go
2. <- 
3. chan

coba cari 3 hal ini
1. \bgob\
2. \b<-\b
3. \b chan\b
*/
func sendEmail(bioChan chan string) {
	fmt.Println("send email!")
	bioChan <- "isSuccessSent=true"
}

func greetHani() {
	fmt.Print("create Progress Bar")
	for idx := 0; idx < 50; idx++ {
		fmt.Print("=")
	}
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
	// buat channel dulu
	bioChannel := make(chan string)
	haniChannel := make(chan string)

	// running secara paralel greetBio dan greetHani
	// hasil yang mau di-return disimpan ke
	go sendEmail(bioChannel)
	go greetHani()

	fmt.Println(<-bioChannel)

	fmt.Println(<-haniChannel)
}
