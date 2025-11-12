package main

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// Copied from demo0, workaround for Ubuntu 24.04
	// Launch browser
	u := launcher.New().NoSandbox(true).MustLaunch()
	browser := rod.New().ControlURL(u).MustConnect()
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
