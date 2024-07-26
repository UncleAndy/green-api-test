package greenapi

import "fmt"

type SendFileByURLRequest struct {
	ChatId          string `json:"chatId"`
	URLFile         string `json:"urlFile"`
	FileName        string `json:"fileName"`
	Caption         string `json:"caption,omitempty"`
	QuotedMessageId string `json:"quotedMessageId,omitempty"`
}

type SendFileByURLResponse struct {
	IDMessage string `json:"idMessage"`
}

func (api *Client) SendFileByURL(id, token string, request SendFileByURLRequest) (*SendFileByURLResponse, error) {
	var response SendFileByURLResponse

	if len(id) < 4 {
		return nil, fmt.Errorf("idInstance is too short")
	}

	// TODO

	return &response, nil
}
