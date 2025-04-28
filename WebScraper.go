package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

type PageInfo struct {
	URL   string
	Title string
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://www.google.com",
		"https://www.github.com",
	}

	var pages []PageInfo

	for _, url := range urls {
		title, err := fetchTitle(url)
		if err != nil {
			fmt.Printf("Error fetching %s: %v\n", url, err)
			continue
		}
		pages = append(pages, PageInfo{URL: url, Title: title})
	}

	fmt.Println("\nScraped Page Titles:")
	for _, page := range pages {
		fmt.Printf("URL: %s\nTitle: %s\n\n", page.URL, page.Title)
	}
}

func fetchTitle(url string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 Chrome/117.0.0.0 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error parsing HTML: %w", err)
	}

	return extractTitle(doc), nil
}

func extractTitle(n *html.Node) string {
	if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
		return n.FirstChild.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		title := extractTitle(c)
		if title != "" {
			return title
		}
	}
	return "No Title Found"
}
