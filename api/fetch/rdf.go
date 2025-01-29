package fetch

import (
	"encoding/xml"
	"fmt"
	"main/models"
	"net/http"
)

// FetchRDFFeed RDFフィードを取得して解析
func FetchRDFFeed(url string) ([]models.RDFItem, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch RDF feed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var rdf models.RDF
	err = xml.NewDecoder(resp.Body).Decode(&rdf)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RDF feed: %w", err)
	}

	return rdf.Items, nil
}