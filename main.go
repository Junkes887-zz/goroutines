package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func manipulateHTML(res io.ReadCloser) string {
	doc, err := goquery.NewDocumentFromReader(res)

	if err != nil {
		fmt.Println("ERROR: manipulate HTML'", err, "'")
	}

	var title string

	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		op, _ := s.Attr("property")
		name, _ := s.Attr("name")
		con, _ := s.Attr("content")

		if name == "twitter:title" || op == "twitter:title" {
			title = con
		}
	})

	return title
}

func scrap(url string, rchan chan string) {
	defer close(rchan)

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: It can't scrap '", url, "'")
	}

	defer res.Body.Close()

	r := manipulateHTML(res.Body)

	if err != nil {
		fmt.Println("ERROR: It can't parse html '", url, "'")
	}

	rchan <- r
}

func scrapListURL(urlToProcess []string) []string {
	var rchan []chan string
	var result []string

	for i, url := range urlToProcess {
		rchan = append(rchan, make(chan string))
		go scrap(url, rchan[i])
	}

	for i := range rchan {
		for r := range rchan[i] {
			result = append(result, r)
		}
	}

	return result
}

func main() {
	urlToProcess := []string{
		"https://medium.freecodecamp.org/how-to-columnize-your-code-to-improve-readability-f1364e2e77ba",
		"https://medium.freecodecamp.org/how-to-think-like-a-programmer-lessons-in-problem-solving-d1d8bf1de7d2",
		"https://medium.freecodecamp.org/code-comments-the-good-the-bad-and-the-ugly-be9cc65fbf83",
		"https://uxdesign.cc/learning-to-code-or-sort-of-will-make-you-a-better-product-designer-e76165bdfc2d",
	}

	var ini time.Time
	ini = time.Now()

	r := scrapListURL(urlToProcess)
	fmt.Println("Without go routine:")

	for _, res := range r {
		fmt.Println(res)
	}
	fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")
}
