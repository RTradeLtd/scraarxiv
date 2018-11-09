package scraarxiv

import (
	"log"
	"os"

	"github.com/RTradeLtd/cmd"
	"github.com/RTradeLtd/config"
)

var (
	// Version denotes the tag of this build
	Version string

	tCfg config.TemporalConfig
)

var commands = map[string]cmd.Cmd{
	"scrap": cmd.Cmd{
		Blurb:       "Starts our arxiv scraper",
		Description: "Scrapes arxiv and takes papers from arxiv, pins it through temporal, and scrapes it",
		Action: func(cfg config.TemporalConfig, args map[string]string) {

		},
	},
}

func main() {
	temporal := cmd.New(commands, cmd.Config{
		Name:     "Scraarxiv",
		ExecName: "scraarxiv",
		Version:  "mvp",
		Desc:     "Scraarxiv pulls data from arvix, stores on ipfs via temporal, and indexes it with Lens",
	})

	if exit := temporal.PreRun(os.Args[1:]); exit == cmd.CodeOK {
		os.Exit(0)
	}

	configDag := os.Getenv("CONFIG_DAG")
	if configDag == "" {
		log.Fatal("CONFIG_DAG is not set")
	}
	tCfg, err := config.LoadConfig(configDag)
	if err != nil {
		log.Fatal(err)
	}
	// load arguments
	flags := make(map[string]string)

	// run our program
	os.Exit(temporal.Run(*tCfg, flags, os.Args[1:]))
}
