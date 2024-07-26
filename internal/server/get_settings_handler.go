package server

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

func (s *Server) GetSettings(ctx *fasthttp.RequestCtx) {
	id := string(ctx.QueryArgs().Peek("idInstance"))
	token := string(ctx.QueryArgs().Peek("apiTokenInstance"))

	ctx.SetContentType("text/plain")

	resp, err := s.client.GetSettings(id, token)
	if err != nil {
		log.Printf("error: %v", err)
		ctx.SetBody([]byte("Error: " + err.Error()))
		ctx.SetStatusCode(fasthttp.StatusOK)
		return
	}

	body, err := json.MarshalIndent(resp, "<br>", "&nbsp;&nbsp;&nbsp;&nbsp;")
	if err != nil {
		log.Printf("error: %v", err)
		ctx.SetBody([]byte("Error: " + err.Error()))
		ctx.SetStatusCode(fasthttp.StatusOK)
		return
	}

	ctx.SetBody(body)
}
