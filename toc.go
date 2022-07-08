package bsmi_go_toc

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"golang.org/x/net/html"
	"strings"
)

type TocItem struct {
	ParentId  string
	ParentLv  int
	ItemValue string
	UniqId    string
	Lv        int
	Children  []TocItem
}

//生成目录html
func Toc_generate(input string) (outStr string, afterHtml string, err error) {
	// Parse the HTML into nodes
	root, e := html.Parse(strings.NewReader(input))
	if e != nil {
		return "", "", e
	}

	// Create and fill the document
	doc := goquery.NewDocumentFromNode(root)

	var rawTocItemList []TocItem
	doc.Find("h1,h2,h3,h4,h5,h6").Each(func(orderIdx int, selection *goquery.Selection) {
		fmt.Println("selection.Nodes[0].Data ->  " + selection.Nodes[0].Data + " title -> " + selection.Text())

		tmpUniqId := Toc_newid()
		rawTocItemList = append(rawTocItemList, TocItem{
			ParentId:  "",
			ParentLv:  0,
			ItemValue: selection.Text(),
			Lv:        getLv(selection.Nodes[0].Data),
			UniqId:    tmpUniqId,
			Children:  nil,
		})

		selection.SetAttr("id", tmpUniqId)
	})

	fmt.Printf(" rawTocItemList: +%v\n", rawTocItemList)

	outStr = generateToc(rawTocItemList)
	ret, _ := doc.Html()
	fmt.Println("doc out html:" + ret[25:len(ret)-14])
	fmt.Println("doc out html end ")
	return outStr, ret[25 : len(ret)-14], nil
}

func generateToc(list []TocItem) string {
	return "<div class=\"bsmi-toc\">" + generateTocHtml(list) + "</div>"

}

func generateTocHtml(list []TocItem) string {
	var outHtml = ""
	var prevLevel = GetMinLevel(list) - 1
	var i = 0
	var ii int
	var h TocItem
	var nextLevel = 0

	listLen := len(list)
	if listLen <= 0 {
		return ""
	}

	outHtml += "<h2>目录</h2>\n"

	for i = 0; i < listLen; i++ {
		h = list[i]

		if (i + 1) < listLen {

			nextLevel = list[i+1].Lv

		} else {
			nextLevel = 0
		}
		if prevLevel == h.Lv {
			outHtml += "<li>\n"
		} else {
			for ii = prevLevel; ii < h.Lv; ii++ {
				outHtml += "<ul><li>\n"
			}
		}

		outHtml += "<a href=\"#" + h.UniqId + "\">" + h.ItemValue + "</a>\n"

		if nextLevel == h.Lv || nextLevel == 0 {
			outHtml += "</li>\n"

			if nextLevel == 0 {
				outHtml += "</ul>\n"
			}
		} else {
			for ii = h.Lv; ii > nextLevel; ii-- {
				if ii == nextLevel+1 {
					outHtml += "</li></ul></li><li>"
				} else {
					outHtml += "</li></ul>"
				}
			}
		}
		prevLevel = h.Lv
	}

	return outHtml
}

func GetMinLevel(list []TocItem) int {
	var i = 9
	var minLevel = 9

	for i = 0; i < len(list); i++ {
		if list[i].Lv < minLevel {
			minLevel = list[i].Lv
		}

		// do not proceed if we have reached absolute minimum
		if minLevel == 1 {
			return minLevel
		}
	}
	return minLevel
}

func getLv(data string) int {
	switch data {
	case "h1":
		return 1
	case "h2":
		return 2
	case "h3":
		return 3
	case "h4":
		return 4
	case "h5":
		return 5
	case "h6":
		return 6
	case "h7":
		return 7
	case "h8":
		return 8
	case "h9":
		return 9
	default:
		return 0

	}
}

func Toc_newid() string {
	id, err := gonanoid.New()

	if err != nil {
		fmt.Println("error new id generate")
	}
	return "btx-id-" + id
}
