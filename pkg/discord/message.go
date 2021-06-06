package discord

import (
	"fmt"

	"github.com/nickname32/discordhook"
)

func Create(data DiscordWebHook) *discordhook.WebhookExecuteParams {
	params := &discordhook.WebhookExecuteParams{}

	info := generateInfo(data)
	params.Embeds = []*discordhook.Embed{
		{
			Title:       data.Title(),
			Description: info,
		},
	}

	return params
}

func generateInfo(data DiscordWebHook) string {
	if data.JobType() == JobStarted {
		return fmt.Sprintf(
			"Host: %s\n"+
				"Job ID: %s (%s)\n",
			data.Host(),
			data.ID(),
			data.URL(),
		)
	}

	return fmt.Sprintf(
		"Duration: %v\n"+
			"Host: %s\n"+
			"Job ID: %s (%s)\n",
		data.Duration(),
		data.Host(),
		data.ID(),
		data.URL(),
	)
}
