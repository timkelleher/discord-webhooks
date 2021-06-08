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
	embed := fmt.Sprintf(
		"Host: %s\n"+
			"Job ID: %s (%s)\n",
		data.Host(),
		data.ID(),
		data.URL(),
	)
	if data.JobType() == JobStarted {
		return embed
	}
	embed = fmt.Sprintf("Duration: %v\n"+embed, data.Duration())

	if data.Success() {
		embed += "<@Tim>\n"
	}

	return embed
}
