package link

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

// Link represents a link (<a href="..."></a>) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse Parse will take in an HTML document and return a slice of links parsed from it.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)

	if err != nil {
		return nil, err
	}

	dfs(doc, "")
	return nil, nil
}

func dfs(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(padding, msg)
	//Iterating through each of the nodes and viewing them. A bit of recursion too. As the loop runs c will be reassigned to the next sibling.
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, padding+"  ")
	}
}
