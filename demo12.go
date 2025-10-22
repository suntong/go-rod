package main

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func main() {
	// Configure browser options
	l := launcher.New().
		Headless(true).
		Devtools(false)

	defer l.Cleanup()
	url := l.MustLaunch()

	browser := rod.New().
		ControlURL(url).
		MustConnect()
	defer browser.MustClose()

	page := browser.MustPage()

	// Navigate to SPA or JS-heavy site
	page.MustNavigate("https://angular.io")

	// Wait for specific element to appear (useful for SPAs)
	page.MustElement("body").MustWaitVisible()
	fmt.Println("Element:body found.")

	// Scroll to load lazy-loaded content
	page.Mouse.MustScroll(0, 1000)
	time.Sleep(2 * time.Second)
	fmt.Println("page scrolled 1000")

	// Click on elements
//	if el, err := page.Element("button.primary"); err == nil {
//		el.MustClick()
//	}

	// Fill forms
//	if el, err := page.Element("input[type='text']"); err == nil {
//		el.MustInput("Hello World")
//	}

	// Extract data after JS rendering
	elements := page.MustElements("h2")
	for _, el := range elements {
		fmt.Println(el.MustText())
	}

	// Execute custom JavaScript
	value := page.MustEval(`() => {
        return {
            url: window.location.href,
            cookies: document.cookie,
            localStorage: Object.keys(localStorage)
        }
    }`).String()
	fmt.Println("JS Result:", value)
}
