package model

import "time"

type News struct {
	Id         int    `json:"id"`
	Title       string `json:"title"`
	Url				 string `json:"url"`
	Description string `json:"description"`
	PublishedAt time.Time `json:"published_at"`
	CreatedAt	 time.Time `json:"created_at"`
	UpdatedAt	 time.Time `json:"updated_at"`
}

type NewsList struct {
	News []News `json:"news"`
}