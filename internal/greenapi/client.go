package greenapi

import "github.com/valyala/fasthttp"

type Client struct {
	client *fasthttp.Client
}

func NewClient() *Client {
	return &Client{
		client: &fasthttp.Client{},
	}
}

func (api *Client) apiUrl(id string) string {
	return "https://" + id[:4] + ".api.greenapi.com"
}
