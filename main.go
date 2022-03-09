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

type Env struct {
	Base_url string
	Endpoint []string
	Port     string
}

type Request struct {
	// The host on which the URL is sought.
	// Per RFC 2616, this is either the value of the Host: header
	// or the host name given in the URL itself.
	// It may be of the form "host:port".
	Host string
}

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func Find(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World"},
	}
	fmt.Println("endpoint hit all article")
	fmt.Printf("Req: %s %s\n", r.Host, r.URL.Path)

	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	env := loadEnvironment()
	found := Find(env.Endpoint, "/")

	if !found {
		fmt.Println("Page not found")
	}

	if found {
		fmt.Fprintf(w, "Welcome to the HomePage!")
	} else {
		w.WriteHeader(404)
	}
}

func stagingPage(w http.ResponseWriter, r *http.Request) {
	env := loadEnvironment()
	found := Find(env.Endpoint, "/staging")
	if !found {
		fmt.Println("Page not found")
	}

	if found {
		fmt.Fprintf(w, "Welcome to the HomePagea!")
	} else {
		w.WriteHeader(404)
	}
}

func loadEnvironment() Env {
	gotenv.Load()
	fmt.Println(os.Getenv("ENVIRONTMENT"))

	configfile := "config/env/" + os.Getenv("ENVIRONTMENT") + ".json"
	envlocal, err := ioutil.ReadFile(configfile)
	if err != nil {
		log.Fatal(err)
	}
	var env Env
	json.Unmarshal([]byte(envlocal), &env)
	return env

}

func main() {
	env := loadEnvironment()

	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/staging", stagingPage)
	log.Fatal(http.ListenAndServe(env.Base_url+":"+env.Port, nil))
}
