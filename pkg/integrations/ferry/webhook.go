package ferry

import "strconv"

type FerryRequest struct {
	Title        string `json:"title"`
	Times        string `json:"times"`
	WebHookID    string `json:"discord_webhook_id"`
	WebHookToken string `json:"discord_webhook_token"`
}

func NewPayload(data FerryRequest) *FerryPayload {
	return &FerryPayload{
		title:        data.Title,
		content:      data.Times,
		webHookID:    data.WebHookID,
		webHookToken: data.WebHookToken,
	}
}

type FerryPayload struct {
	title        string
	content      string
	webHookID    string
	webHookToken string
}

func (fp FerryPayload) Integration() string {
	return "Ferry"
}

func (fp FerryPayload) Title() string {
	return fp.title
}

func (fp FerryPayload) Content() string {
	return fp.content
}

func (fp FerryPayload) WebHookID() int {
	id, err := strconv.Atoi(fp.webHookID)
	if err != nil {
		return 0
	}
	return id
}

func (fp FerryPayload) WebHookToken() string {
	return fp.webHookToken
}
