package main

import (
	"fmt"
	"context"
	"../xscrape"
)

func main() {
	site := "http://www.google.com"
	sitemap, err := xscrape.Site(context.Background(), site)
	fmt.Println(err)
	fmt.Println(sitemap)
}
