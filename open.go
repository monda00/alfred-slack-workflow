package main

import (
	aw "github.com/deanishe/awgo"
)

func openChannel() {
	c := aw.NewCache(cache_dir)

	var c_channels []Channel
	if c.Exists(cache_file) {
		if err := c.LoadJSON(cache_file, &c_channels); err != nil {
			wf.FatalError(err)
		}

		for _, channel := range c_channels {
			wf.NewItem(channel.Name).
				Var("teamID", channel.TeamID).
				Var("channelID", channel.ID).
				Valid(true)
		}
	}

	args := wf.Args()
	if len(args) > 1 {
		wf.Filter(args[1])
	}

	wf.SendFeedback()
}
