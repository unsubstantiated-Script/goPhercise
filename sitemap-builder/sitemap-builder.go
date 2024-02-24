package sitemap_builder

import (
	"flag"
	"fmt"
	link "goPhercise/html-link-parser"
	"io"
	"net/http"
	"net/url"
	"strings"
)

/*
	1. GET the webpage
	2. Parse all the links on the page
	3. Build proper urls with our links
	4. Filter out any links w/ a diff domain
	5. Find all the pages (Breadth First Search)
	6. Print out XML
*/

func RollSitemapBuilder() {
	//Getting the URL
	urlFlag := flag.String("url", "https://gophercises.com", "The url you wanna build a sitemap for.")
	flag.Parse()

	fmt.Println(*urlFlag)

	pages := get(*urlFlag)
	//pages := hrefs(resp.Body, base)

	for _, page := range pages {
		fmt.Println(page)
	}

}

func get(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	return hrefs(resp.Body, base)
}

func hrefs(r io.Reader, base string) []string {
	links, _ := link.Parse(r)

	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}

	return ret
}
