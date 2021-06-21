package discord

import (
	"context"
	"fmt"

	"github.com/andersfylling/snowflake"
	"github.com/nickname32/discordhook"
)

type GenericWebHook interface {
	Integration() string
	Title() string
	Content() string
}

//https://pkg.go.dev/github.com/nickname32/discordhook
func NewGenericWebHook(conf DiscordConfig, data GenericWebHook) {
	wa, err := discordhook.NewWebhookAPI(snowflake.Snowflake(conf.WebHookID()), conf.WebHookToken(), true, nil)
	if err != nil {
		panic(err)
	}

	ctx := context.Context(context.Background())

	wh, err := wa.Get(ctx)
	if err != nil {
		fmt.Println("Failed to fetch webhook")
		panic(err)
	}
	fmt.Println("Webhook identified:", wh.Name)

	params := &discordhook.WebhookExecuteParams{}
	params.Embeds = []*discordhook.Embed{
		{
			Title:       data.Title(),
			Description: data.Content(),
		},
	}

	msg, err := wa.Execute(ctx, params, nil, "")
	if err != nil {
		fmt.Println("Failed to publish message:")
		panic(err)
	}

	fmt.Println("Message published:", msg.ID)
}
