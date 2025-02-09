package models

type RDFItem struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
	PublishedAt string    `xml:"date"`
}

type RDF struct {
	Items []RDFItem `xml:"item"`
}