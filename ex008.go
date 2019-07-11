package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
<?xml version = "1.0" encoding = "utf-8"?>
<!-- xslplane.1.xml -->
<?xml-stylesheet type = "text/xsl"  href = "xslplane.1.xsl" ?>
<plane>
   <year> 1977 </year>
   <make> Cessna </make>
   <model> Skyhawk </model>
   <color> Light blue and white </color>
</plane>

*/

type Plane struct {
	Year int `xml:"year"`
	Make string `xml:"make"`
	ModelName string `xml:"model"`
	ColorType string `xml:"color"`
}

func main() {
	resp, _ := http.Get("https://www.cs.utexas.edu/~mitra/csFall2015/cs329/lectures/xml/xslplanes.1.xml.txt")
	bytes, _ := ioutil.ReadAll(resp.Body)
	stringBody := string(bytes)
	fmt.Println(stringBody)
	_ = resp.Body.Close()

	var xmlData Plane
	_ = xml.Unmarshal(bytes, &xmlData)

	fmt.Println(xmlData.Year)
	fmt.Println(xmlData.Make)
	fmt.Println(xmlData.ModelName)
	fmt.Println(xmlData.ColorType)
}