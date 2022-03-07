package main

import (
	"encoding/json"
	"fmt"
)

// type Bird struct {
// 	Species     string
// 	Description string
// }

type Env struct {
	Base_url string
	Endpoint string
	Port     string
}

func main() {
	// content, err := ioutil.ReadFile("config/env/local.json")
	// var data interface{}
	// fmt.Printf("Type: %T, value: %v", content, content)
	// json.Unmarshal(content, &data)

	// fmt.Printf("File contents: %s", content)
	// fmt.Printf("File contents: %s", content)

	// birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	// var bird Bird
	// json.Unmarshal([]byte(birdJson), &bird)
	// fmt.Printf("Species: %s, Description: %s ", bird.Species, bird.Description)

	envlocal := `{"base_url": "localhost","endpoint": "bla","port": "8080"}`
	var el Env
	json.Unmarshal([]byte(envlocal), &el)
	fmt.Printf("base url: %s, endpoint: %s, port: %s ", el.Base_url, el.Endpoint, el.Port)

	// data := []byte(`[{"href":"/publication/192a7f47-84c1-445e-a615-ff82d92e2eaa/article/b;version=1493756861347"},{"href":"/publication/192a7f47-84c1-445e-a615-ff82d92e2eaa/article/a;version=1493756856398"}]`)

	// fmt.Printf("%T %v \n \n", data, data)
	// fmt.Printf("%T %v \n \n", content, content)
	// fmt.Printf("%T %v \n \n", err, err)
	// // data, _ := ioutil.ReadFile("config/env/local.json")
	// var objmap []map[string]interface{}
	// if err := json.Unmarshal(content, &objmap); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(objmap[0]["port"]) // to parse out your value

}
