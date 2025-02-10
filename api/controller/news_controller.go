package controller

import (
	"fmt"
	"main/usecase"
)

type INewsController interface {
	FetchAndStoreFeeds(urls []string)
}

type newsController struct {
	newsUsecase usecase.INewsUsecase
}

// NewNewsController は NewsController のインスタンスを作成
func NewNewsController(newsUsecase usecase.INewsUsecase) INewsController {
	return &newsController{newsUsecase}
}

// FetchAndStoreFeeds はフィードを取得してDBに保存する
func (nc *newsController) FetchAndStoreFeeds(urls []string) {
	for _, url := range urls {
		fmt.Printf("Fetching feed from: %s\n", url)

		// フィードを取得
		newsItems, err := nc.newsUsecase.ParseFeed(url)
		if err != nil {
			fmt.Printf("Error fetching feed from %s: %v\n", url, err)
			continue
		}

		// 取得したニュースをDBに保存
		for _, news := range newsItems {
			fmt.Printf("- %s\n  Link: %s\n  Description: %s\n\n", news.Title, news.Url, news.Description)

			if err := nc.newsUsecase.InsertNews(news); err != nil {
				fmt.Printf("Error inserting news into DB: %v\n", err)
			}
		}
	}
}