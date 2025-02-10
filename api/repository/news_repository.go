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
	query := `
    INSERT INTO news (title, url, description, published_at, created_at, updated_at)
    VALUES ($1, $2, $3, $4, NOW(), NOW())
	`
	_, err := r.db.Exec(query, news.Title, news.Url, news.Description, news.PublishedAt)
	return err
}