package main

import (
  "fmt"
  "testing"
  "github.com/gocolly/colly"
)

func TestBasicColly(t *testing.T) {
  basic_sample("http://go-colly.org/")
}
