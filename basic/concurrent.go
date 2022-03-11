package basic

import "fmt"

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
