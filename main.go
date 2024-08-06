package main

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/gocolly/colly/v2"

	"parliamenthack-api/type/model"
	"parliamenthack-api/util/log"
	"parliamenthack-api/util/value"
)

func main() {
	// * Instantiate default collector
	c := colly.NewCollector()

	// * Construct items array
	var items []*model.ScrapedItem

	// * Iterate each search result
	c.OnHTML("div.xcR77", func(e *colly.HTMLElement) {
		// * Extract image url
		var imageUrl *string
		e.ForEachWithBreak(".oR27Gd", func(i int, element *colly.HTMLElement) bool {
			imageUrl = value.Ptr(element.ChildAttr("img", "src"))
			return false
		})

		// * Extracting item title
		title := e.ChildText(".rgHvZc > a")

		// * Extracting item price
		price := e.ChildText(".HRLxBb")

		// * Filter out empty items
		if imageUrl == nil || title == "" || price == "" {
			return
		}

		// * Construct scraped item
		item := &model.ScrapedItem{
			ImageUrl: imageUrl,
			Name:     &title,
			Price:    &price,
		}

		// * Append to items array
		items = append(items, item)
	})

	// Start the collector
	searchQuery := "ทีวีจอแบน 32 นิ้ว ซัมซุง"
	searchURL := "https://www.google.com/search?tbm=shop&q=" + strings.ReplaceAll(searchQuery, " ", "+")

	log.Debug("Visiting search URL", searchURL)
	err := c.Visit(searchURL)
	if err != nil {
		log.Fatal("Failed to visit search URL", err)
	}

	spew.Dump(items)
}
