package cli

import (
	"time"

	"github.com/timkelleher.com/discord-webhooks/pkg/discord"
)

func NewPayload(msg string, id int, token string) *DevPayload {
	return &DevPayload{
		Msg:          msg,
		PayloadID:    id,
		PayloadToken: token,
		Succeeded:    true,
	}
}

type DevPayload struct {
	Msg          string
	PayloadID    int
	PayloadToken string
	Succeeded    bool
}

func (dp DevPayload) ID() string {
	return "1"
}

func (dp DevPayload) JobType() discord.JobType {
	return discord.JobCompleted
}

func (dp DevPayload) Title() string {
	return dp.Msg
}

func (dp DevPayload) Success() bool {
	return dp.Succeeded
}

func (dp DevPayload) TimeStart() string {
	return "start"
}

func (dp DevPayload) TimeEnd() string {
	return "end"
}

func (dp DevPayload) Duration() time.Duration {
	return 1 * time.Second
}

func (dp DevPayload) URL() string {
	return ""
}

func (dp DevPayload) Source() string {
	return "source"
}

func (dp DevPayload) Host() string {
	return "host"
}

func (dp DevPayload) WebHookID() int {
	return dp.PayloadID
}

func (dp DevPayload) WebHookToken() string {
	return dp.PayloadToken
}
