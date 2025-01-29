package main

import (
	"fmt"
	"log"

	"main/config"
	"main/db"
	"main/fetch"
)

func main() {
	// 環境変数のロード
	config.LoadEnv()

	// DB接続
	database, err := db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer database.Close()

	// RSS/RDF URLリスト
	urls := []string{
		"https://www.nhk.or.jp/rss/news/cat4.xml",
		"https://assets.wor.jp/rss/rdf/sankei/politics.rdf",
		"https://assets.wor.jp/rss/rdf/yomiuri/politics.rdf",
		"https://assets.wor.jp/rss/rdf/ynnews/politics.rdf",
	}

	// フィード取得とデータ挿入
	for _, url := range urls {
		fmt.Printf("Fetching feed from: %s\n", url)

		if url[len(url)-4:] == ".rdf" {
			items, err := fetch.FetchRDFFeed(url)
			if err != nil {
				fmt.Printf("Error fetching RDF feed from %s: %v\n", url, err)
				continue
			}

			for _, item := range items {
				fmt.Printf("- %s\n  Link: %s\n  Description: %s\n\n", item.Title, item.Link, item.Desc)
				err = db.InsertNews(database, item.Title, item.Link, item.Desc)
				if err != nil {
					fmt.Printf("Error inserting RDF item into DB: %v\n", err)
				}
			}
		} else {
			items, err := fetch.FetchRSSFeed(url)
			if err != nil {
				fmt.Printf("Error fetching RSS feed from %s: %v\n", url, err)
				continue
			}

			for _, item := range items {
				fmt.Printf("- %s\n  Link: %s\n  Description: %s\n\n", item.Title, item.Link, item.Description)
				err = db.InsertNews(database, item.Title, item.Link, item.Description)
				if err != nil {
					fmt.Printf("Error inserting RSS item into DB: %v\n", err)
				}
			}
		}
	}
}