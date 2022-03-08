package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"github.com/subosito/gotenv"
)

type Eenv struct {
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

func init() {
	gotenv.Load()
}

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
	var env Eenv
	json.Unmarshal([]byte(envlocal), &env)

	fmt.Println(os.Getenv("ENVIRONTMENT"))

	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(env.Base_url+":"+env.Port, nil))
}

func main() {
	log.Println(os.Getenv("APP_ID"))
	handleRequests()
}
