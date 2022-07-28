package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/foods", httpHandler)
	fmt.Println("run on :8080")
	http.ListenAndServe(":8080", nil)
}

func httpHandler(w http.ResponseWriter, req *http.Request) {

	response, err := getProducts()

	if err != nil {

		fmt.Fprintf(w, err.Error()+"\r\n")

	} else {

		enc := json.NewEncoder(w)
		enc.SetIndent("", "  ")

		if err := enc.Encode(response); err != nil {
			fmt.Println(err.Error())
		}

	}
}
