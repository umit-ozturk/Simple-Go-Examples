package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request)  {
	articles := Articles{
		Article{Title:"Test Title", Desc: "Test Description", Content: "Test Content"},
	}

	fmt.Println("Endpoint: Home All Articles")
	json.NewEncoder(w).Encode(articles)
	
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Home Endpoint Hit.")
}

func handleRequest()  {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main()  {
	handleRequest()
}