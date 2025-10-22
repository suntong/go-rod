package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
)

type Repository struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Stars       string `json:"stars"`
	Language    string `json:"language"`
}

func main() {
	// Create stealth browser (bypass detection)
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// Create page with stealth mode
	page := stealth.MustPage(browser)

	// Set viewport
	page.MustSetViewport(1920, 1080, 1, false)

	// Set extra headers
	page.MustSetExtraHeaders("Accept-Language", "en-US,en;q=0.9")

	// Navigate with timeout
	err := page.Timeout(30 * time.Second).Navigate("https://github.com/trending")
	if err != nil {
		log.Fatal(err)
	}

	// Wait for content to load
	page.MustWaitLoad()

	// Wait for specific elements (for SPAs)
	page.MustElement("article").MustWaitVisible()

	// Scroll to load all content
	page.MustEval(`() => {
        return new Promise((resolve) => {
            let totalHeight = 0;
            let distance = 100;
            let timer = setInterval(() => {
                let scrollHeight = document.body.scrollHeight;
                window.scrollBy(0, distance);
                totalHeight += distance;

                if(totalHeight >= scrollHeight){
                    clearInterval(timer);
                    resolve();
                }
            }, 100);
        });
    }`)

	// Extract data using JavaScript
	data := page.MustEval(`() => {
        const articles = Array.from(document.querySelectorAll('article'));
        return articles.map(article => ({
            title: article.querySelector('h2 a').innerText.trim(),
            description: article.querySelector('p')?.innerText.trim() || '',
            url: article.querySelector('h2 a').href,
            stars: article.querySelector('a[href$="/stargazers"]').innerText.trim(),
            language: article.querySelector('span[itemprop="programmingLanguage"]')?.innerText.trim() || ''
        }));
    }`)

	// Parse and output as JSON
	var repositories []Repository
	if err := data.Unmarshal(&repositories); err != nil {
		log.Fatal(err)
	}

	jsonOutput, err := json.MarshalIndent(repositories, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonOutput))
}
