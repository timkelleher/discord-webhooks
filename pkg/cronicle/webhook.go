package cronicle

import (
	"fmt"
	"strconv"
	"time"

	"github.com/timkelleher.com/discord-webhooks/pkg/discord"
)

func NewPayload(data WebHookRequest) *WebHookPayload {
	return &WebHookPayload{payload: data}
}

//https://github.com/jhuckaby/Cronicle#event-web-hook
type WebHookRequest struct {
	ID        string            `json:"id"`
	JobType   string            `json:"action"`
	Title     string            `json:"event_title"`
	Code      int               `json:"code"`
	Elapsed   int               `json:"elapsed"`
	TimeStart string            `json:"time_start"`
	TimeEnd   string            `json:"time_end"`
	URL       string            `json:"job_details_url"`
	Source    string            `json:"source"`
	Host      string            `json:"hostname"`
	Params    map[string]string `json:"params"`
}

type WebHookPayload struct {
	payload WebHookRequest
}

func (whp WebHookPayload) ID() string {
	return whp.payload.ID
}

func (whp WebHookPayload) JobType() discord.JobType {
	if whp.payload.JobType == "job_start" {
		return discord.JobStarted
	}
	return discord.JobCompleted
}

func (whp WebHookPayload) Title() string {
	if whp.payload.JobType == "job_start" {
		return "Start: " + whp.payload.Title
	}

	var result string
	if whp.Success() {
		result = "Success"
	} else {
		result = "Failed"
	}
	return fmt.Sprintf("Completed: %s (%s)", whp.payload.Title, result)
}

func (whp WebHookPayload) Success() bool {
	return whp.payload.Code == 0
}

func (whp WebHookPayload) TimeStart() string {
	return whp.payload.TimeStart
}

func (whp WebHookPayload) TimeEnd() string {
	return whp.payload.TimeEnd
}

func (whp WebHookPayload) Duration() time.Duration {
	return time.Duration(whp.payload.Elapsed) * time.Second
}

func (whp WebHookPayload) URL() string {
	return whp.payload.URL
}

func (whp WebHookPayload) Source() string {
	return whp.payload.Source
}

func (whp WebHookPayload) Host() string {
	return whp.payload.Host
}

func (whp WebHookPayload) WebHookID() int {
	id, err := strconv.Atoi(whp.payload.Params["discord_webhook_id"])
	if err != nil {
		return 0
	}
	return id
}

func (whp WebHookPayload) WebHookToken() string {
	return whp.payload.Params["discord_webhook_token"]
}
