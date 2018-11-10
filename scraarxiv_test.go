package scraarxiv_test

import (
	"fmt"
	"testing"

	"github.com/RTradeLtd/config"
	"github.com/RTradeLtd/scraarxiv/magnifier"
	"github.com/RTradeLtd/scraarxiv/searcher"
)

const (
	testConfigPath = "test/config.json"
	testURL1       = "http://arxiv.org/pdf/1711.03577v1"
	testURL2       = "http://arxiv.org/pdf/1801.07883v2"
	testURL3       = "http://arxiv.org/pdf/1703.02910v1"
	testURL4       = "http://arxiv.org/pdf/1805.03551v2"
	testURL5       = "http://arxiv.org/pdf/1708.05866v2"
	testURL6       = "http://arxiv.org/pdf/1710.06798v1"
	testURL7       = "http://arxiv.org/pdf/1801.00631v1"
)

var (
	urls = []string{testURL1, testURL2, testURL3, testURL4, testURL4, testURL5, testURL6, testURL7}
)

func TestScraarxivNoSearch(t *testing.T) {
	t.Skip()
	cfg, err := config.LoadConfig(testConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	glass, err := magnifier.NewGlassClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	// test the download
	files, err := glass.DownloadFiles(urls)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(files)
}

func TestScraarxivSearch(t *testing.T) {
	fmt.Println("loading config")
	cfg, err := config.LoadConfig(testConfigPath)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("loading glass client")
	glass, err := magnifier.NewGlassClient(cfg)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("running searcher")
	urls, err := searcher.Search("deep learning", 1)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("extracting pdf urls")
	pdfURLs := searcher.ExtractPDFURLs(urls)
	if len(pdfURLs) == 0 {
		t.Fatal("failed to get pdf urls")
	}
	fmt.Println("manigfying files")
	if err = glass.Magnify(pdfURLs); err != nil {
		t.Fatal(err)
	}
}
