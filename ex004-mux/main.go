package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "Home Endpoint Hit.")
}

func handleRequest()  {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8001", nil))
}

func main()  {
	handleRequest()
}