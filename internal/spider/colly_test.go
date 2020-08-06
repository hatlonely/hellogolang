package spider

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/gocolly/colly/v2"
)

func TestColly(t *testing.T) {
	c := colly.NewCollector(
		colly.Async(),
		colly.TraceHTTP(),
		colly.URLFilters(regexp.MustCompile(`http://go-colly.org/*`)),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL, r.ID, r.Depth)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(string(r.Body))
	})

	c.Visit("http://go-colly.org/")

	c.Wait()
}
