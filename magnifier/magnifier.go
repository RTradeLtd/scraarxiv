// Package magnifier is used to take a pdf from arxiv, download it, store on IPFS, and request for it to be indexed by Lens
package magnifier

import (
	"context"
	"fmt"
	"net/http"

	"github.com/RTradeLtd/config"
	ipfsapi "github.com/RTradeLtd/go-ipfs-api"
	pb "github.com/RTradeLtd/grpc/lens/request"
	"github.com/RTradeLtd/scraarxiv/lens"
)

// Glass is used to take an arxiv pdf, put it into IPFS, and index it with Lens
type Glass struct {
	s *ipfsapi.Shell
	l *lens.Client
}

// NewGlassClient instantiates our Glass client to magnify arxiv papers
func NewGlassClient(cfg *config.TemporalConfig) (*Glass, error) {
	shell := ipfsapi.NewShell(fmt.Sprintf("%s:%s", cfg.IPFS.APIConnection.Host, cfg.IPFS.APIConnection.Port))
	// check our connection
	if _, err := shell.ID(); err != nil {
		return nil, err
	}
	lensClient, err := lens.NewClient(cfg.Endpoints)
	if err != nil {
		return nil, err
	}
	return &Glass{
		s: shell,
		l: lensClient,
	}, nil
}

// Magnify is used to take a PDF urls, download it, inject into Temporal pin system, and index with Lens
func (g *Glass) Magnify(urls []string, maxDownloads int) error {
	fmt.Println("downloading files")
	var hashes []string
	for _, v := range urls {
		resp, err := http.Get(v)
		if err != nil {
			continue
		}
		hash, err := g.s.Add(resp.Body)
		if err != nil {
			resp.Body.Close()
			continue
		}
		resp.Body.Close()
		hashes = append(hashes, hash)
	}
	fmt.Println("indexing content in lens")
	// index the content with Lens
	for _, v := range hashes {
		if _, err := g.l.SubmitIndexRequest(
			context.Background(),
			&pb.IndexRequest{
				DataType:         "ipld",
				ObjectIdentifier: v},
		); err != nil {
			fmt.Println("error encountered processing ", err)
			continue
		}
	}
	//TODO: pin the content
	fmt.Println("hashes ", hashes)
	return nil
}
