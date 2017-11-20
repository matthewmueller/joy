package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

var domInterfaces = []string{
	"Attr",
	"CharacterData",
	"ChildNode ",
	"Comment",
	"CustomEvent",
	"Document",
	"DocumentFragment",
	"DocumentType",
	"DOMError",
	"DOMException",
	"DOMImplementation",
	"DOMString",
	"DOMTimeStamp",
	"DOMSettableTokenList",
	"DOMStringList",
	"DOMTokenList",
	"Element",
	"Event",
	"EventTarget",
	"HTMLCollection",
	"MutationObserver",
	"MutationRecord",
	"NamedNodeMap",
	"Node",
	"NodeFilter",
	"NodeIterator",
	"NodeList",
	"NonDocumentTypeChildNode",
	"ParentNode",
	"ProcessingInstruction",
	"Selection",
	"Range",
	"Text",
	"TextDecoder ",
	"TextEncoder ",
	"TimeRanges",
	"TreeWalker",
	"URL",
	"Window",
	"Worker",
	"XMLDocument ",
}

func ExampleScrape() {
	doc, err := goquery.NewDocument("https://developer.mozilla.org/en-US/docs/Web/API/Attr")
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("#Properties + dl > dt").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Find("a").Text()
		// title := s.Find("i").Text()
		fmt.Printf("Review %d: %s\n", i, band)
	})
}

func main() {
	ExampleScrape()
}
