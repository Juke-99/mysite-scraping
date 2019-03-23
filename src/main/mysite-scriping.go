package main

import (
  "github.com/gocolly/colly"

	"encoding/json"
	"io/ioutil"
  "fmt"
  "bytes"
)

type Index struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func main() {
  var d []string

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("li a", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	for _, index := range jsonRead("src/main/json/mysite-index.json") {
		c.Visit(index.Url)
    fmt.Println("---------")
	}
}

func jsonRead(jsonUri string) ([]Index) {
	byteArray, err := ioutil.ReadFile(jsonUri)

	if err != nil {
		fmt.Println(err)
	}

	var index []Index

	if err := json.Unmarshal(byteArray, &index); err != nil {
		fmt.Println(err)
	}

	return index
}

func createJson() {
  // data := new(Data)

  // data := map[string]interface{}{
  //   "data": map[string]interface{}{
  //     "url":
  //   }
  // }

  jsonData, err := json.Marshal(data)

  if err != nil {
    fmt.Println(err)
    return
  }

  output := new(bytes.Buffer)
  json.Indent(output, jsonData, "", "  ")
  fmt.Println(output.String())
}
