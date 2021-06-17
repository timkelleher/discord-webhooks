package discord

import (
	"context"
	"fmt"
	"time"

	"github.com/andersfylling/snowflake"
	"github.com/nickname32/discordhook"
)

type JobType int

const (
	JobStarted JobType = iota
	JobCompleted
)

type DiscordWebHook interface {
	Integration() string
	ID() string
	JobType() JobType
	Title() string
	Success() bool
	TimeStart() string
	TimeEnd() string
	Duration() time.Duration
	URL() string
	Source() string
	Host() string
}

type DiscordConfig interface {
	WebHookID() int
	WebHookToken() string
}

//https://pkg.go.dev/github.com/nickname32/discordhook
func NewWebHook(conf DiscordConfig, data DiscordWebHook) {
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

	msgParams := Create(data)

	// Some jobs are so fast, Cronicle seems to publish complete events
	// before start events, or it's batched or something.
	if data.JobType() == JobCompleted {
		time.Sleep(1 * time.Second)
	}

	msg, err := wa.Execute(ctx, msgParams, nil, "")
	if err != nil {
		fmt.Println("Failed to publish message:")
		panic(err)
	}

	fmt.Println("Message published:", msg.ID)
}
