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

func (api *Client) SendMessage(id, token string, request SendMessageRequest) (*SendMessageResponse, error) { // nolint: dupl
	var response SendMessageResponse

	if len(id) < 4 {
		return nil, fmt.Errorf("idInstance is too short")
	}

	url := api.apiUrl(id) + "/waInstance" + id + "/sendMessage/" + token

	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req := fasthttp.AcquireRequest()
	defer req.Reset()
	req.Header.SetContentType("application/json")
	req.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.SetBody(reqBody)

	resp := fasthttp.AcquireResponse()
	defer resp.Reset()
	err = api.client.Do(req, resp)
	if err != nil {
		return nil, err
	}

	status := resp.StatusCode()
	if status != fasthttp.StatusOK {
		return nil, fmt.Errorf("unexpected status %d", status)
	}

	body := resp.Body()

	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
