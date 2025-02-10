package model

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PublishedAt string    `xml:"pubDate"`
}

type RSSChannel struct {
	Items []RSSItem `xml:"item"`
}

type RSS struct {
	Channel RSSChannel `xml:"channel"`
}