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

	//Get Request
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	//These defers should be kept close to their OG instantiations. They will run at the end regardless. A "return" won't shut them down.

	//Even though this gets pushed to the end, it's still best to keep it here.
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

	// The url from the response
	reqUrl := resp.Request.URL

	// Setting up the base url via the url struct from GoLang
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}

	//converting to string with some Go magic
	base := baseUrl.String()

	pages := hrefs(resp.Body, base)

	for _, page := range pages {
		fmt.Println(page)
	}

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
