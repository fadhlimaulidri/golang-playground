package main

import (
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Article []Article

func allArticle(w http.ResponseWriter)  {
	
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home page endpoint hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
