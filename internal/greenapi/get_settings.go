package greenapi

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type GetSettingsResponse struct {
	WID                               string `json:"wid"`
	CountryInstance                   string `json:"countryInstance,omitempty"`
	TypeAccount                       string `json:"typeAccount,omitempty"`
	WebhookUrl                        string `json:"webhookUrl,omitempty"`
	WebhookUrlToken                   string `json:"webhookUrlToken"`
	DelaySendMessagesMilliseconds     int    `json:"delaySendMessagesMilliseconds"`
	MarkIncomingMessagesReaded        string `json:"markIncomingMessagesReaded"`
	MarkIncomingMessagesReadedOnReply string `json:"markIncomingMessagesReadedOnReply"`
	SharedSession                     string `json:"sharedSession,omitempty"`
	OutgoingWebhook                   string `json:"outgoingWebhook"`
	OutgoingMessageWebhook            string `json:"outgoingMessageWebhook"`
	OutgoingAPIMessageWebhook         string `json:"outgoingAPIMessageWebhook"`
	IncomingWebhook                   string `json:"incomingWebhook"`
	DeviceWebhook                     string `json:"deviceWebhook,omitempty"`
	StatusInstanceWebhook             string `json:"statusInstanceWebhook,omitempty"`
	StateWebhook                      string `json:"stateWebhook"`
	EnableMessagesHistory             string `json:"enableMessagesHistory,omitempty"`
	KeepOnlineStatus                  string `json:"keepOnlineStatus"`
	PollMessageWebhook                string `json:"pollMessageWebhook"`
	IncomingBlockWebhook              string `json:"incomingBlockWebhook,omitempty"`
	IncomingCallWebhook               string `json:"incomingCallWebhook"`
}

func (api *Client) GetSettings(id, token string) (*GetSettingsResponse, error) {
	var response GetSettingsResponse

	if len(id) < 4 {
		return nil, fmt.Errorf("idInstance is too short")
	}

	url := api.apiUrl(id) + "/waInstance" + id + "/getSettings/" + token

	status, body, err := api.client.Get(nil, url)
	if err != nil {
		return nil, err
	}
	if status != fasthttp.StatusOK {
		return nil, fmt.Errorf("unexpected status %d", status)
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
