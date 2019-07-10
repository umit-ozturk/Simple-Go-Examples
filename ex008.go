package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("https://www.cs.utexas.edu/~mitra/csFall2015/cs329/lectures/xml/xslplanes.1.xml.txt")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	_ = resp.Body.Close()
}