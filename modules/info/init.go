package info

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/diraven/sugo"
	"github.com/dustin/go-humanize"
	"gitlab.com/diraven/crabot/utils"
	"runtime"
	"time"
)

// Init initializes module on the given bot.
func Init(sg *sugo.Instance) {
	sg.AddCommand(cmd)
}

// Info shows some general bot info.
var cmd = &sugo.Command{
	Trigger:     "info",
	Description: "Shows basic info about bot.",
	Execute: func(req *sugo.Request) (resp *sugo.Response, err error) {
		// Set command response.
		now := time.Now().UTC()

		// Get memory usage stats.
		var memStats runtime.MemStats
		runtime.ReadMemStats(&memStats)

		_, err = req.Sugo.Session.ChannelMessageSendEmbed(req.Channel.ID, &discordgo.MessageEmbed{
			URL:         "https://github.com/diraven/sugo",
			Title:       "https://github.com/diraven/sugo",
			Description: "A Discord bot written in Go.",
			Timestamp:   utils.TimeToDiscordTimestamp(now),
			Color:       sugo.ColorInfo,
			Footer: &discordgo.MessageEmbedFooter{
				Text: "DiRaven#0519",
			},
			Fields: []*discordgo.MessageEmbedField{
				{
					Name:   "sugo ver.",
					Value:  sugo.VERSION,
					Inline: true,
				},
				{
					Name:   "discordgo ver.",
					Value:  discordgo.VERSION,
					Inline: true,
				},
				{
					Name:   "go ver.",
					Value:  runtime.Version(),
					Inline: true,
				},
				{
					Name:   "memory usage",
					Value:  fmt.Sprintf("%s", humanize.Bytes(memStats.Sys)),
					Inline: true,
				},
				{
					Name:   "goroutines",
					Value:  fmt.Sprintf("%d", runtime.NumGoroutine()),
					Inline: true,
				},
				{
					Name:   "modules loaded",
					Value:  fmt.Sprintf("%d", len(req.Sugo.RootCommand.SubCommands)),
					Inline: true,
				},
			},
		})
		if err != nil {
			return
		}
		return
	},
}