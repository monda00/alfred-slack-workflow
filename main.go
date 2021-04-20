package main

import (
	"flag"

	aw "github.com/deanishe/awgo"
)

var (
	wf         *aw.Workflow
	cache_dir  = "./cache"
	cache_file = "cache.json"
)

type Channel struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	TeamID string `json:"teamid"`
}

func init() {
	wf = aw.New()
}

func run() {
	update := flag.Bool("update", false, "Update Channels")
	open := flag.Bool("open", false, "Open Channel")
	flag.Parse()

	if *update {
		updateChannels()
		return
	}

	if *open {
		openChannel()
		return
	}

}

func main() {
	wf.Run(run)
}
