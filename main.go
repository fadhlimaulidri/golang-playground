package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Env struct {
	Base_url string
	Endpoint string
	Port     string
}

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}
	fmt.Println("endpoint hit all article")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	// envlocal := `{"base_url": "localhost","endpoint": "bla","port": ":8080"}`
	envlocal, err := ioutil.ReadFile("config/env/local.json")
	if err != nil {
		log.Fatal(err)
	}
	var el Env
	json.Unmarshal([]byte(envlocal), &el)
	// fmt.Printf("base url: %s, endpoint: %s, port: %s ", el.Base_url, el.Endpoint, el.Port)

	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(el.Port, nil))
}

func main() {
	// Articles = []Article{
	// 	Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
	// 	Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	// }

	handleRequests()

}
