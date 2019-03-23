// package main
//
// import (
//   "github.com/gocolly/colly"
//
//   "fmt"
// )
//
// func main() {
// 	c := colly.NewCollector()
//
// 	// Find and visit all links
// 	c.OnHTML("li a", func(e *colly.HTMLElement) {
// 		fmt.Println(e.Text)
// 	})
//
// 	c.Visit("https://www.juke-99.com/blog/engineer/okio-buffer.html")
//   c.Visit("https://www.juke-99.com/blog/profile.html")
//   c.Visit("https://www.juke-99.com/blog/engineer/okhttp-mockwebserver.html")
//   c.Visit("https://www.juke-99.com/blog/engineer/stream-in-java.html")
// }
