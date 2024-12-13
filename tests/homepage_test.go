package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/playwright-community/playwright-go"
)

func TestHomepage(t *testing.T) {
	fmt.Println(time.Now())
	// List of browsers to test
	browserNames := []string{"chromium", "firefox", "webkit"}

	// Iterate over each browser
	for _, browserName := range browserNames {
		t.Run(browserName, func(t *testing.T) {
			err := testHomepageForBrowser(browserName)
			if err != nil {
				t.Errorf("Test failed for browser %s: %v", browserName, err)
			}
		})
	}
	fmt.Println(time.Now())
}

func testHomepageForBrowser(browserName string) error {
	// Initialize Playwright and launch the browser
	pw, err := playwright.Run()
	if err != nil {
		return fmt.Errorf("could not start playwright: %w", err)
	}
	defer pw.Stop()

	var browser playwright.Browser
	switch browserName {
	case "chromium":
		browser, err = pw.Chromium.Launch()
	case "firefox":
		browser, err = pw.Firefox.Launch()
	case "webkit":
		browser, err = pw.WebKit.Launch()
	default:
		return fmt.Errorf("unsupported browser: %s", browserName)
	}

	if err != nil {
		return fmt.Errorf("could not launch %s: %w", browserName, err)
	}
	defer browser.Close()

	// Open a new page and navigate to the homepage
	context, err := browser.NewContext()
	if err != nil {
		return fmt.Errorf("could not create context: %w", err)
	}
	defer context.Close()

	page, err := context.NewPage()
	if err != nil {
		return fmt.Errorf("could not create page: %w", err)
	}

	// Navigate to the homepage
	url := "https://www.wasserstrom.com"
	fmt.Println("request begin")
	fmt.Println(time.Now())
	if _, err := page.Goto(url); err != nil {
		return fmt.Errorf("could not navigate to %s: %w", url, err)
	}
	fmt.Println(time.Now())

	title, err := page.Title()
	if err != nil {
		return fmt.Errorf("could not get page title: %w", err)
	}
	if title == "" {
		return fmt.Errorf("empty title for the page: %s", title)
	}

	return nil
}
