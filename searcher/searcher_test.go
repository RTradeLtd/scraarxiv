package searcher_test

import (
	"fmt"
	"testing"

	"github.com/RTradeLtd/scraarxiv/searcher"
)

func TestSearcher(t *testing.T) {
	urls, err := searcher.Search("deep learning", 1)
	if err != nil {
		t.Fatal(err)
	}
	pdfURLs := searcher.ExtractPDFURLs(urls)
	fmt.Println(pdfURLs)
}
