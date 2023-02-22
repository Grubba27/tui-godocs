package main

import (
	"fmt"
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/charmbracelet/glamour"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	url := os.Args[1]
	c := &http.Client{
		Timeout: 10 * time.Second,
	}
	res, err := c.Get(url)
	if err != nil {
		log.Fatal("err")
	}

	body, err := goquery.NewDocumentFromReader(res.Body)
	err = res.Body.Close()
	if err != nil {
		log.Fatal("err")
	}
	html, _ := body.Find("article").Html()
	converter := md.NewConverter("", true, nil)
	mkd, err := converter.ConvertString(html)

	lines := strings.Split(mkd, "\n")
	slachCount := 0
	var markdown []string
	for _, line := range lines {
		if strings.Contains(line, "```") {
			slachCount++
			if slachCount%2 != 0 {
				line = line + "go"
			}
		}
		markdown = append(markdown, line)
	}
	m := strings.Join(markdown, "\n")
	out, _ := glamour.Render(m, "dracula")
	fmt.Print(out)
}
