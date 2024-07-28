package server

import (
	"encoding/json"
	"log"

	"github.com/UncleAndy/green-api-test/internal/greenapi"
	"github.com/valyala/fasthttp"
)

func (s *Server) SendMessage(ctx *fasthttp.RequestCtx) {
	id := string(ctx.QueryArgs().Peek("idInstance"))
	token := string(ctx.QueryArgs().Peek("apiTokenInstance"))
	chatId := string(ctx.QueryArgs().Peek("chatIdSend"))
	msg := string(ctx.QueryArgs().Peek("message"))

	ctx.SetContentType("text/plain")

	resp, err := s.client.SendMessage(id, token, greenapi.SendMessageRequest{
		ChatId:  chatId,
		Message: msg,
	})
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
