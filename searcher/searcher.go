// Package searcher is used to search for research papers on arxiv
package searcher

import (
	"context"
	"fmt"

	arxiv "github.com/orijtech/arxiv/v1"
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
