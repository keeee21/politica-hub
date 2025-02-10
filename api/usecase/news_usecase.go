package usecase

import (
	"encoding/xml"
	"fmt"
	"io"
	"main/model"
	"main/repository"
	"main/util"
	"net/http"
	"strings"
	"time"
)

type INewsUsecase interface {
	ParseFeed(url string) ([]model.News, error)
	InsertNews(news model.News) error
}

type newsUsecase struct {
	newsRepo repository.INewsRepository
}

func NewNewsUsecase(newsRepo repository.INewsRepository) INewsUsecase {
	return &newsUsecase{newsRepo}
}

// ParseFeed は指定された URL からフィードをパースする
func (nu *newsUsecase) ParseFeed(url string) ([]model.News, error) {
	data, contentType, err := fetchFeed(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch feed: %w", err)
	}

	// RDF or RSS を判定
	if strings.Contains(contentType, "rdf+xml") {
		return parseRDF(data)
	}
	return parseRSS(data)
}

// fetchFeed は指定された URL からフィードデータを取得する
func fetchFeed(url string) ([]byte, string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, "", fmt.Errorf("failed to fetch feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", fmt.Errorf("failed to read response body: %w", err)
	}

	return data, resp.Header.Get("Content-Type"), nil
}

// parseRDF は RDF フィードをパースする
func parseRDF(data []byte) ([]model.News, error) {
	var rdf model.RDF
	if err := xml.Unmarshal(data, &rdf); err != nil {
		return nil, fmt.Errorf("failed to parse RDF feed: %w", err)
	}

	var newsItems []model.News
	for _, item := range rdf.Items {
		pubTime, err := util.ParsePublishedAt(item.PublishedAt)
		if err != nil {
			fmt.Printf("Warning: failed to parse date '%s', using current time\n", item.PublishedAt)
			pubTime = time.Now()
		}

		newsItems = append(newsItems, model.News{
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Desc,
			PublishedAt: pubTime,
		})
	}

	return newsItems, nil
}

// parseRSS は RSS フィードをパースする
func parseRSS(data []byte) ([]model.News, error) {
	var rss model.RSS
	if err := xml.Unmarshal(data, &rss); err != nil {
		return nil, fmt.Errorf("failed to parse RSS feed: %w", err)
	}

	var newsItems []model.News
	for _, item := range rss.Channel.Items {
		pubTime, err := util.ParsePublishedAt(item.PublishedAt)
		if err != nil {
			fmt.Printf("Warning: failed to parse date '%s', using current time\n", item.PublishedAt)
			pubTime = time.Now()
		}

		newsItems = append(newsItems, model.News{
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubTime,
		})
	}

	return newsItems, nil
}

// InsertNews はニュースをデータベースに挿入する
func (nu *newsUsecase) InsertNews(news model.News) error {
	return nu.newsRepo.InsertNews(&news)
}