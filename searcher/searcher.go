// Package searcher is used to search for research papers on arxiv
package searcher

import (
	"context"
	"fmt"
	"strings"

	arxiv "github.com/orijtech/arxiv/v1"
)

const (
	basePDFURL = "https://arxiv.org/pdf"
)

// Search is used to perform a search against arxiv
func Search(term string, maxPageNumbers int64) ([]string, error) {
	var urlsToScrape []string
	// construct our query and generate a channel to receive data one
	responseChannel, cancel, err := arxiv.Search(
		context.Background(),
		&arxiv.Query{
			Terms:         term,
			MaxPageNumber: maxPageNumbers,
		})

	if err != nil {
		return nil, err
	}

	for page := range responseChannel {
		// if this page had an error, skip it
		if err = page.Err; err != nil {
			fmt.Printf("error occured: %s\n", err)
			continue
		}
		for _, entry := range page.Feed.Entry {
			urlsToScrape = append(urlsToScrape, entry.ID)
		}
		if page.PageNumber > 5 {
			cancel()
		}
	}
	return urlsToScrape, nil
}

// ExtractPDFURLs is used to take an arxiv paper url, and get its pdf download equivalent
func ExtractPDFURLs(urls []string) []string {
	var pdfURLs []string
	for _, v := range urls {
		split := strings.Split(v, "/")
		url := fmt.Sprintf("%s/%s", basePDFURL, split[len(split)-1])
		pdfURLs = append(pdfURLs, url)
	}
	return pdfURLs
}
