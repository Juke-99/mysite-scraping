// package main
//
// import (
// 	"strings"
//
// 	"golang.org/x/net/html"
// 	"fmt"
// )
//
// func main() {
// 	// url := "https://www.juke-99.com/blog/engineer/okio-buffer.html"
//
// 	// res, err := http.Get(url)
//
// 	// if err != nil {
// 	// 	panic(err)
// 	// }
//
// 	// defer res.Body.Close()
//
// 	// byteArray, err := ioutil.ReadAll(res.Body)
// 	// fmt.Println(string(byteArray))
//
// 	//s := string(byteArray)
// 	s := `<aside class="category-box"><h3 class="category-header">カテゴリー</h3><ul><li><a>java</a></li><li><a>Okio v1</a></li><li><a>Buffer</a></li></ul></aside>`
// 	// fmt.Print(n.Data)
// 	z := html.NewTokenizer(strings.NewReader(s))
//
// 	for {
// 		tt := z.Next()
//
// 		if tt == html.StartTagToken {
// 			name, attr := z.TagName()
// 			fmt.Println(string(name))
// 			fmt.Println(attr)
// 			fmt.Println(string(z.Text()))
//
// 			if tt == html.EndTagToken {
// 				break
// 			}
// 		}
// 	}
// }
