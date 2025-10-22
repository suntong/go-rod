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
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser)
	page.MustSetViewport(1920, 1080, 1, false)
	page.MustSetExtraHeaders("Accept-Language", "en-US,en;q=0.9")

	var allRepositories []Repository

	for i := 1; i <= 38; i++ {
		url := fmt.Sprintf("https://gitcode.com/search?q=%%E5%%89%%8D%%E5%%90%%8E%%E5%%88%%86%%E7%%A6%%BB+%%E5%%BF%%AB%%E9%%80%%9F%%E5%%BC%%80%%E5%%8F%%91+%%E5%%B9%%B3%%E5%%8F%%B0&type=repo&l=Go&p=%d", i)
		fmt.Printf("Scraping page %d: %s\n", i, url)

		err := page.Timeout(30 * time.Second).Navigate(url)
		if err != nil {
			log.Printf("Failed to navigate to page %d: %v", i, err)
			continue
		}
		page.MustWaitLoad()

		// Wait for the search results to be visible
		page.MustElement(".search-result-item").MustWaitVisible()

		data := page.MustEval(`() => {
			const items = Array.from(document.querySelectorAll('.search-result-item'));
			return items.map(item => {
				const titleElement = item.querySelector('.title a');
				const descElement = item.querySelector('.description');
				const starsElement = item.querySelector('.stars');
				const langElement = item.querySelector('.language');

				return {
					title: titleElement ? titleElement.innerText.trim() : '',
					description: descElement ? descElement.innerText.trim() : '',
					url: titleElement ? titleElement.href : '',
					stars: starsElement ? starsElement.innerText.trim() : '',
					language: langElement ? langElement.innerText.trim() : ''
				};
			});
		}`)

		var repositories []Repository
		if err := data.Unmarshal(&repositories); err != nil {
			log.Printf("Failed to unmarshal data from page %d: %v", i, err)
			continue
		}
		allRepositories = append(allRepositories, repositories...)
		time.Sleep(1 * time.Second) // Be polite
	}

	jsonOutput, err := json.MarshalIndent(allRepositories, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonOutput))
}
