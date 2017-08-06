package main

import (
	"log"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://golang.org/ref/spec")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	removeID("nav", doc)
	removeID("topbar", doc)
	removeID("playground", doc)
	removeID("lowframe", doc)

	file, err := os.Create("spec.html")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	if err = html.Render(file, doc); err != nil {
		log.Fatal(err)
	}
}

func removeID(id string, node *html.Node) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		var found = false
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, a := range n.Attr {
				if a.Key == "id" && a.Val == id {
					n.Parent.RemoveChild(n)
					found = true
					break
				}
			}
		}
		if found {
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node)
}
