package db

import (
	"database/sql"
	"fmt"
	"main/config"
	"time"

	_ "github.com/lib/pq"
)

// ConnectDB PostgreSQLに接続
func ConnectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.GetEnv("DB_HOST"),
		config.GetEnv("DB_PORT"),
		config.GetEnv("DB_USER"),
		config.GetEnv("DB_PASSWORD"),
		config.GetEnv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to DB: %w", err)
	}
	return db, nil
}

// InsertNews ニュースデータをDBに挿入
func InsertNews(db *sql.DB, title, url, description string, publishedAt time.Time) error {
	query := `
        INSERT INTO news (title, url, description, published_at)
        VALUES ($1, $2, $3, $4)
        ON CONFLICT (url) DO NOTHING;
    `
	_, err := db.Exec(query, title, url, description, publishedAt)
	return err
}