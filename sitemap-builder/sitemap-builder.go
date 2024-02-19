package sitemap_builder

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
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
	io.Copy(os.Stdout, resp.Body)
}
