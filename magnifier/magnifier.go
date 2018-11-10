// Package magnifier is used to take a pdf from arxiv, download it, store on IPFS, and request for it to be indexed by Lens
package magnifier

import (
	"fmt"

	"github.com/RTradeLtd/config"
	ipfsapi "github.com/RTradeLtd/go-ipfs-api"
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
func (g *Glass) Magnify(urls []string) error {
	return nil
}
