package main

import (
	"fmt"

	"github.com/go-rod/rod"
)

func main() {
	// Launch browser
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// Create a new page
	page := browser.MustPage("https://example.com")

	// Wait for JavaScript to load
	page.MustWaitLoad()

	// Execute JavaScript
	result := page.MustEval(`() => {
        return document.title;
    }`)
	fmt.Println("Page title:", result)

	// Take screenshot
	page.MustScreenshot("screenshot.png")

	// Get HTML after JS rendering
	html := page.MustHTML()
	fmt.Println("Rendered HTML:", html)
}
