package greenapi

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type SendMessageRequest struct {
	ChatId          string `json:"chatId"`
	Message         string `json:"message"`
	QuotedMessageId string `json:"quotedMessageId,omitempty"`
	LinkPreview     bool   `json:"linkPreview,omitempty"`
}

type SendMessageResponse struct {
	IDMessage string `json:"idMessage"`
}

func (api *Client) SendMessage(id, token string, request SendMessageRequest) (*SendMessageResponse, error) {
	var response SendMessageResponse

	if len(id) < 4 {
		return nil, fmt.Errorf("idInstance is too short")
	}

	url := api.apiUrl(id) + "/waInstance" + id + "/sendMessage/" + token

	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	status, body, err := api.client.Post(reqBody, url, nil)
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
