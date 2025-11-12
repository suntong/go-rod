package main

import (
	"fmt"
	"github.com/go-rod/rod"
)

func main() {
	// Create a data URL with simple HTML
	html := `
		<html>
			<body>
				<h1>Test Page</h1>
				<div id="content">
					<p>This is paragraph one.</p>
					<p>This is paragraph two.</p>
				</div>
			</body>
		</html>
	`

	page := rod.New().MustConnect().MustPage("data:text/html," + html)

	// Select the div element
	el := page.MustElement("#content")

	// 1. Get HTML (Outer)
	inner, _ := el.HTML()
	fmt.Println("--- HTML ---")
	fmt.Println(inner)
	fmt.Println("--------------------")

	// 2. Get Outer HTML
	prop, _ := el.Property("outerHTML")
	fmt.Println("--- Outer HTML ---")
	fmt.Println(prop.Str()) // Use .Str() to get the string value
	fmt.Println("--------------------")

// --- Inner HTML via Explicit Property ---
// This explicitly gets the "innerHTML" property.
innerProp, _ := el.Property("innerHTML")
fmt.Println("--- Inner HTML (from Property) ---")
fmt.Println(innerProp.Str())
fmt.Println("----------------------------------")
}
