package cronicle

import (
	"fmt"
	"strconv"
	"time"

	"github.com/timkelleher.com/discord-webhooks/pkg/discord"
)

func NewPayload(data CronicleRequest) *CroniclePayload {
	return &CroniclePayload{payload: data}
}

//https://github.com/jhuckaby/Cronicle#event-web-hook
type CronicleRequest struct {
	ID        string            `json:"id"`
	JobType   string            `json:"action"`
	Title     string            `json:"event_title"`
	Code      int               `json:"code"`
	Elapsed   float32           `json:"elapsed"`
	TimeStart string            `json:"time_start"`
	TimeEnd   string            `json:"time_end"`
	URL       string            `json:"job_details_url"`
	Source    string            `json:"source"`
	Host      string            `json:"hostname"`
	Params    map[string]string `json:"params"`
}

type CroniclePayload struct {
	payload CronicleRequest
}

func (whp CroniclePayload) Integration() string {
	return "Cronicle"
}

func (whp CroniclePayload) ID() string {
	return whp.payload.ID
}

func (whp CroniclePayload) JobType() discord.JobType {
	if whp.payload.JobType == "job_start" {
		return discord.JobStarted
	}
	return discord.JobCompleted
}

func (whp CroniclePayload) Title() string {
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

func (whp CroniclePayload) Success() bool {
	return whp.payload.Code == 0
}

func (whp CroniclePayload) TimeStart() string {
	return whp.payload.TimeStart
}

func (whp CroniclePayload) TimeEnd() string {
	return whp.payload.TimeEnd
}

func (whp CroniclePayload) Duration() time.Duration {
	seconds := time.Duration(whp.payload.Elapsed)
	return seconds * time.Second
}

func (whp CroniclePayload) URL() string {
	return whp.payload.URL
}

func (whp CroniclePayload) Source() string {
	return whp.payload.Source
}

func (whp CroniclePayload) Host() string {
	return whp.payload.Host
}

func (whp CroniclePayload) WebHookID() int {
	id, err := strconv.Atoi(whp.payload.Params["discord_webhook_id"])
	if err != nil {
		return 0
	}
	return id
}

func (whp CroniclePayload) WebHookToken() string {
	return whp.payload.Params["discord_webhook_token"]
}
