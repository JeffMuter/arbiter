package tests

import (
	"github.com/playwright-community/playwright-go"
	"testing"
)

func TestHomepage(t *testing.T) {
	// Setup browser and page
	pw, err := playwright.Run()
	if err != nil {
		t.Fatalf("could not start Playwright: %v", err)
	}
	defer pw.Stop()

}
