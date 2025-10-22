package main

import (
	"fmt"
	"github.com/go-rod/rod"
)

func main() {
	page := rod.New().MustConnect().MustPage("https://angular.dev/")

	// This is the corrected code
	elements := page.MustElements("h2")
	for _, el := range elements {
		fmt.Println(el.MustText())
	}
}
