package searcher_test

import (
	"testing"

	"github.com/RTradeLtd/scraarxiv/searcher"
)

func TestSearcher(t *testing.T) {
	if err := searcher.Search("deep learning", 5); err != nil {
		t.Fatal(err)
	}
}
