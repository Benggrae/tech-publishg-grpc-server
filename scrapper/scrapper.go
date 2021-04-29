package scrapper

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type wohaTechDoc struct {
	author string
	time   time.Time
	title  string
	link   string
}

func getMonthInt(month string) int {
	if strings.LastIndex(month, "Jan") > -1 {
		return 1
	}
	if strings.LastIndex(month, "Feb") > -1 {
		return 2
	}
	if strings.LastIndex(month, "Mar") > -1 {
		return 3
	}
	if strings.LastIndex(month, "Apr") > -1 {
		return 4
	}
	if strings.LastIndex(month, "May") > -1 {
		return 5
	}
	if strings.LastIndex(month, "Jun") > -1 {
		return 6
	}
	if strings.LastIndex(month, "Jul") > -1 {
		return 7
	}
	if strings.LastIndex(month, "Aug") > -1 {
		return 8
	}
	if strings.LastIndex(month, "Sep") > -1 {
		return 9
	}
	if strings.LastIndex(month, "Oct") > -1 {
		return 10
	}
	if strings.LastIndex(month, "Nov") > -1 {
		return 11
	}
	if strings.LastIndex(month, "Dec") > -1 {
		return 12
	}
	return -1

}

const Wowha = "Wowha"

func WoowaScrapper() {
	//컬리 생성
	c := colly.NewCollector()
	woowaRoot := "https://woowabros.github.io"
	defer c.Visit(woowaRoot)
	var woahDocList []wohaTechDoc

	//html 찾기
	c.OnHTML(".list", func(elList *colly.HTMLElement) {
		elList.ForEach(".list-module", func(index int, el *colly.HTMLElement) {

			doc := wohaTechDoc{}
			postMeta := el.ChildText(".post-meta")

			meta := strings.Split(postMeta, ",")

			doc.author = strings.TrimSpace(meta[len(meta)-1])
			year, _ := strconv.Atoi(strings.TrimSpace(meta[1]))
			tempText := strings.Split(meta[0], " ")
			month := getMonthInt(tempText[0])
			day, _ := strconv.Atoi(tempText[1])

			doc.time = time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
			doc.link = el.ChildAttr("a", "href")

			woahDocList = append(woahDocList, doc) //배열추가

			//println(year, month, day)

			//println(strings.Split(meta[0], " ")[0])

		})
	})

	//저장후
	c.OnScraped(func(r *colly.Response) {
		fmt.Print(woahDocList)
	})

	defer fmt.Println(woahDocList)

	c.OnRequest(func(r *colly.Request) {
		log.Println("방문 ::", r.URL)
	})

}
