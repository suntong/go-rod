package main

import (
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

// This example demonstrates how to fill out and submit a form.
func main() {
	page := rod.New().
		MustConnect().
		Trace(true). // log useful info about what rod is doing
		Timeout(15 * time.Second).
		MustPage("https://github.com/search")

	page.MustElement(`input[data-component=input]`).MustWaitVisible().MustInput("rod").MustType(input.Enter)

	res := page.MustElementR("a", "rod").MustParent().MustParent().MustParent().MustNext().MustText()

	log.Printf("got: `%s`", strings.TrimSpace(res))
	// got: `A Chrome DevTools Protocol driver for web automation and scraping.`
}
