package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/stealth"
)

type ScrapedData struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Links       []string `json:"links"`
	Images      []string `json:"images"`
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
        return {
            title: document.title,
            description: document.querySelector('meta[name="description"]')?.content || '',
            links: Array.from(document.querySelectorAll('a')).map(a => a.href).slice(0, 50),
            images: Array.from(document.querySelectorAll('img')).map(img => img.src).slice(0, 20)
        }
    }`)

	// Parse result
	var scraped ScrapedData
	jsonStr := data.String()
	json.Unmarshal([]byte(jsonStr), &scraped)

	fmt.Printf("Title: %s\n", scraped.Title)
	fmt.Printf("Found %d links and %d images\n", len(scraped.Links), len(scraped.Images))

	// Handle cookies
	cookies := page.MustCookies()
	fmt.Printf("Cookies: %v\n", cookies)

	// Take full page screenshot
	page.MustScreenshotFullPage("fullpage.png")
}
