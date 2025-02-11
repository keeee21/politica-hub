package main

import (
	"log"
	"main/controller"

	"github.com/robfig/cron/v3"
)

// StartBatchScheduler はバッチ処理をスケジューリングする
func StartBatchScheduler(newsController controller.INewsController) {
	// RSS/RDF URLリスト
	urls := []string{
		"https://www.nhk.or.jp/rss/news/cat4.xml",
		"https://assets.wor.jp/rss/rdf/sankei/politics.rdf",
		"https://assets.wor.jp/rss/rdf/yomiuri/politics.rdf",
		"https://assets.wor.jp/rss/rdf/ynnews/politics.rdf",
	}

	// cron スケジューラーを作成
	c := cron.New()

	// 30分毎 に実行
	c.AddFunc("*/30 * * * *", func() {
		log.Println("Running batch job...")
		newsController.FetchAndStoreFeeds(urls)
		log.Println("Batch job completed.")
	})

	// cron を開始
	c.Start()

	log.Println("Batch scheduler started...")

	// プログラムが終了しないように無限ループ
	select {}
}
