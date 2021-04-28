package scrapper

import (
	"log"
	"time"

	"github.com/gocolly/colly"
)

type wohaTechDoc struct {
	author string
	time   time.Time
	title  string
	link   string
}

func RunScrapper() {
	c := colly.NewCollector()

	defer c.Visit("https://woowabros.github.io/")
	c.OnHTML(".list", func(elList *colly.HTMLElement) {
		elList.ForEach(".list-module", func(index int, el *colly.HTMLElement) {
			doc := wohaTechDoc{}
			doc.author = el.ChildText(".post-meta")

		})
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("방문 ::", r.URL)
	})

}
