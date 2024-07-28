package server

import (
	"encoding/json"
	"log"
	"path"
	"strings"

	"github.com/UncleAndy/green-api-test/internal/greenapi"
	"github.com/valyala/fasthttp"
)

func (s *Server) SendFileByURL(ctx *fasthttp.RequestCtx) {
	id := string(ctx.QueryArgs().Peek("idInstance"))
	token := string(ctx.QueryArgs().Peek("apiTokenInstance"))
	chatId := string(ctx.QueryArgs().Peek("chatIdFile"))
	urlFile := string(ctx.QueryArgs().Peek("urlFile"))

	ctx.SetContentType("text/plain")

	fileName := strings.TrimSpace(path.Base(urlFile))
	if len(fileName) == 0 {
		fileName = "no-name"
	}

	resp, err := s.client.SendFileByURL(id, token, greenapi.SendFileByURLRequest{
		ChatId:   chatId,
		URLFile:  urlFile,
		FileName: fileName,
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
