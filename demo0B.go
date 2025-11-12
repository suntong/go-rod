package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-rod/rod"
)

var (
	pURL  = flag.String("url", "https://example.com", "URL to test")
	pWait = flag.String("w", "a", "Page element to wait for")
)

func main() {
	flag.Parse()

	url := *pURL
	// url can be passed with -url or as the 1st arg
	if len(flag.Args()) >= 1 {
		url = flag.Args()[0]
	}

	// Launch browser
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	// Create a new page
	fmt.Fprintf(os.Stderr, "Visiting '%s' (& wait for '%s') ...\n", url, *pWait)
	page := browser.MustPage(url)

	// Wait for JavaScript to load
	page.MustWaitLoad()

	// Explicit Wait for it to load
	page.MustElement(*pWait).MustWaitVisible()

	// Get body HTM
	html := page.MustElement("body").MustHTML()
	fmt.Println(html)
}

// go run demo0B.go -url='https://so.gitee.com/?q=Vue%20go' -w 'a[aria-label=Next]'
// go run demo0B.go -w 'a[aria-label=Next]' 'https://so.gitee.com/?q=Vue%20go'
