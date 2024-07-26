package server

import "github.com/valyala/fasthttp"

const (
	indexTmplName = "index.tmpl"
)

func (s *Server) Index(ctx *fasthttp.RequestCtx) {
	s.renderTemplate(ctx, indexTmplName, nil)

	ctx.SetContentType("text/html")
}
