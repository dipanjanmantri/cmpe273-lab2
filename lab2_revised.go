package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"net/http"
)

type PostResponse struct {
	Greeting string
}

type ReqJson struct {
	person string
}

func hello_get(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func hello_post(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	decoder := json.NewDecoder(req.Body)
	var reqJson ReqJson
	err := decoder.Decode(&reqJson)
	if err != nil {
		panic("Error decoding json.")
	}
	post_response := PostResponse{Greeting: "Hello, " + reqJson.person + "!"}
	json.NewEncoder(rw).Encode(post_response)
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello_get)
	mux.POST("/hello/", hello_post)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}