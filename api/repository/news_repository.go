package repository

import (
	"database/sql"
	"main/model"
)

// INewsRepository はニュースデータを扱うリポジトリのインターフェース
type INewsRepository interface {
	InsertNews(news *model.News) error
}

// newsRepository は NewsRepository の実装
type newsRepository struct {
	db *sql.DB
}

// NewNewsRepository は newsRepository のインスタンスを作成
func NewNewsRepository(db *sql.DB) INewsRepository {
	return &newsRepository{db}
}

// InsertNews はニュースをデータベースに保存する
func (r *newsRepository) InsertNews(news *model.News) error {
	// まず URL がすでに存在するかチェック
	var exists bool
	err := r.db.QueryRow("SELECT EXISTS(SELECT 1 FROM news WHERE url = $1)", news.Url).Scan(&exists)
	if err != nil {
		return err
	}

	// 既に存在する場合は何もしない
	if exists {
		return nil
	}

	// 存在しない場合のみ INSERT
	query := `
    INSERT INTO news (title, url, description, published_at, created_at, updated_at)
    VALUES ($1, $2, $3, $4, NOW(), NOW())
	`
	_, err = r.db.Exec(query, news.Title, news.Url, news.Description, news.PublishedAt)
	return err
}