package searcher_test

import (
	"fmt"
	"testing"

	"github.com/RTradeLtd/scraarxiv/searcher"
)

func TestSearcher(t *testing.T) {
	urls, err := searcher.Search("deep learning", 5)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(urls)
}
