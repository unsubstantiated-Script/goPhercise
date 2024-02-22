package sitemap_builder

import (
	"flag"
	"fmt"
	link "goPhercise/html-link-parser"
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
	urlFlag := flag.String("url", "https://gophercises.com", "The url you wanna build a sitemap for.")
	flag.Parse()

	fmt.Println(*urlFlag)

	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	//These defers should be kept close to their OG instantiations. They will run at the end regardless. A "return" won't shut them down.
	defer resp.Body.Close()

	//Copies from a reader os.Stdout to a writer resp.Body
	//io.Copy(os.Stdout, resp.Body)

	/*
		/some-path
		https://gophercises.com/some-path
		#fragment
			i.e.
			/some-path#fragment
			https://gophercises.com/some-path#fragment
		mailto:jon@calhoun.io
	*/

	reqUrl := resp.Request.URL

	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	base := baseUrl.String()

	links, _ := link.Parse(resp.Body)
	//for _, l := range links {
	//	fmt.Println(l)
	//}

	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	for _, href := range hrefs {
		fmt.Println(href)
	}

}
