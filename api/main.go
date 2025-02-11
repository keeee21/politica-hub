package main

import (
	"log"
	"main/config"
	"main/controller"
	"main/db"
	"main/repository"
	"main/usecase"
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

	// DI (依存関係の注入)
	newsRepo := repository.NewNewsRepository(database)
	newsUsecase := usecase.NewNewsUsecase(newsRepo)
	newsController := controller.NewNewsController(newsUsecase)

	// バッチスケジューラーを起動
	StartBatchScheduler(newsController)
}
