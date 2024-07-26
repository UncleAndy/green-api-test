package greenapi

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
)

type GetStateInstanceResponse struct {
	StateInstance string `json:"stateInstance"`
}

func (api *Client) GetStateInstance(id, token string) (*GetStateInstanceResponse, error) {
	var response GetStateInstanceResponse

	if len(id) < 4 {
		return nil, fmt.Errorf("idInstance is too short")
	}

	url := api.apiUrl(id) + "/waInstance" + id + "/getStateInstance/" + token

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
