package models

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
}

type RSSChannel struct {
	Items []RSSItem `xml:"item"`
}

type RSS struct {
	Channel RSSChannel `xml:"channel"`
}