package fetch

import (
	"encoding/xml"
	"fmt"
	"main/models"
	"net/http"
)

// FetchRSSFeed RSSフィードを取得して解析
func FetchRSSFeed(url string) ([]models.RSSItem, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RSS feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var rss models.RSS
	err = xml.NewDecoder(resp.Body).Decode(&rss)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSS feed: %w", err)
	}

	return rss.Channel.Items, nil
}