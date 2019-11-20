package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("dolarhoje.com"),
	)

	c.OnHTML("input[id=nacional]", func(e *colly.HTMLElement) {
		fmt.Println("Cotação do dólar hoje: " + e.Attr("value"))
	})

	c.Visit("https://dolarhoje.com/")
}
