package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

type Result struct {
	userName string
	title    string
	likes    string
}

func scrap(url string) (r chan Result) {
	res, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: It can't scrap '", url, "'")
	}

	defer res.Body.Close()

	body := res.Body

	htmlParsed, err := html.Parse(body)
	if err != nil {
		fmt.Println("ERROR: It can't parse html '", url, "'")
	}
}

func main() {
	urlToProcess := []string{
		"https://medium.freecodecamp.org/how-to-columnize-your-code-to-improve-readability-f1364e2e77ba",
		"https://medium.freecodecamp.org/how-to-think-like-a-programmer-lessons-in-problem-solving-d1d8bf1de7d2",
		"https://medium.freecodecamp.org/code-comments-the-good-the-bad-and-the-ugly-be9cc65fbf83",
		"https://uxdesign.cc/learning-to-code-or-sort-of-will-make-you-a-better-product-designer-e76165bdfc2d",
	}

	var ini time.Time
	fmt.Println("Without go routine:")
	ini = time.Now()
	for _, url := range urlToProcess {
		r := scrap(url)
		fmt.Println(r)
	}

	fmt.Println("(Took ", time.Since(ini).Seconds(), "secs)")
}
