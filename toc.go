package bsmi_go_toc

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/net/html"
	"strings"
)

func Toc_generate(input string) (string, string, error) {
	// Parse the HTML into nodes
	root, e := html.Parse(strings.NewReader(input))
	if e != nil {
		return "", "", e
	}

	// Create and fill the document
	doc := goquery.NewDocumentFromNode(root)
	var outStr string
	doc.Find("h1,h2,h3,h4,h5,h6").Each(func(orderIdx int, selection *goquery.Selection) {
		fmt.Println("selection.Nodes[0].Data -> " + selection.Nodes[0].Data)
		outStr += selection.Text()
		selection.SetAttr("id", Toc_newid())
	})
	ret, _ := doc.Html()
	fmt.Println("doc out html:" + ret[25:len(ret)-14])
	fmt.Println("doc out html end ")
	return outStr, ret[25 : len(ret)-14], nil
}

func Toc_newid() string {
	id, err := gonanoid.New()

	if err != nil {
		fmt.Println("error new id generate")
	}
	return "btx-id-" + id
}
