package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
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

func testPostArticle(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Test Article Post Endpoint")
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Home Endpoint Hit.")
}

func handleRequest()  {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	myRouter.HandleFunc("/articles", testPostArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", myRouter))
}

func main()  {
	handleRequest()
}