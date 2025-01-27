package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// RSS Item represents a single RSS item
type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

// RSS Channel represents the RSS channel
type RSSChannel struct {
	Items []RSSItem `xml:"item"`
}

// RSS Feed represents the RSS feed
type RSS struct {
	Channel RSSChannel `xml:"channel"`
}

// RDF Item represents a single RDF item
type RDFItem struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
}

// RDF represents the RDF feed
type RDF struct {
	Items []RDFItem `xml:"item"`
}

// fetchRSSFeed fetches and parses an RSS feed
func fetchRSSFeed(url string) (*RSS, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var rss RSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSS feed: %w", err)
	}

	return &rss, nil
}

// fetchRDFFeed fetches and parses an RDF feed
func fetchRDFFeed(url string) (*RDF, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RDF feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var rdf RDF
	err = xml.NewDecoder(resp.Body).Decode(&rdf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RDF feed: %w", err)
	}

	return &rdf, nil
}

func main() {
	urls := []string{
		"https://www.nhk.or.jp/rss/news/cat4.xml",
		"https://assets.wor.jp/rss/rdf/sankei/politics.rdf",
        "https://assets.wor.jp/rss/rdf/yomiuri/politics.rdf",
        "https://assets.wor.jp/rss/rdf/ynnews/politics.rdf",
	}

	for _, url := range urls {
		fmt.Printf("Fetching feed from: %s\n", url)

		if url[len(url)-4:] == ".rdf" {
			// Try RDF parsing
			rdf, err := fetchRDFFeed(url)
			if err != nil {
				fmt.Printf("Error fetching RDF feed from %s: %v\n", url, err)
				continue
			}

			fmt.Println("RDF Feed Items:")
			for i, item := range rdf.Items {
				fmt.Printf("%d. %s\n", i+1, item.Title)
				fmt.Printf("   Link: %s\n", item.Link)
				fmt.Printf("   Description: %s\n\n", item.Desc)
			}
		} else {
			// Try RSS parsing
			rss, err := fetchRSSFeed(url)
			if err != nil {
				fmt.Printf("Error fetching RSS feed from %s: %v\n", url, err)
				continue
			}

			fmt.Println("RSS Feed Items:")
			for i, item := range rss.Channel.Items {
				fmt.Printf("%d. %s\n", i+1, item.Title)
				fmt.Printf("   Link: %s\n", item.Link)
				fmt.Printf("   Description: %s\n\n", item.Description)
			}
		}
	}
}