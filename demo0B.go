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

	// Get body HTM
	html := page.MustElement("body").MustHTML()
	fmt.Println("Body HTML:", html)
}
