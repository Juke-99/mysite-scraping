package main

import (
	"fmt"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

type Document struct {
	*Selection
	Url      *url.URL
	rootNode *html.Node
}

type Selection struct {
	Nodes    []*html.Node
	document *Document
	prevSel  *Selection
}

func main() {
	url := "***"

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	// byteArray, err := ioutil.ReadAll(res.Body)
	root, err := html.Parse(res.Body)

	d := &Document{nil, nil, root}
	d.Selection = &Selection{[]*html.Node{root}, d, nil}

	//fmt.Println(string(byteArray))
	// fmt.Println(string(root))
	fmt.Println(d)
}
