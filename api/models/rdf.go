package models

type RDFItem struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
}

type RDF struct {
	Items []RDFItem `xml:"item"`
}