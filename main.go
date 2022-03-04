package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page endpoint hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	fmt.Println("listen port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
