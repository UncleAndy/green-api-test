package server

import (
	"fmt"
	"html/template"
	"log"
	"path/filepath"
	"strings"

	"github.com/UncleAndy/green-api-test/internal/greenapi"
	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"

	"github.com/UncleAndy/green-api-test/internal/config"
	"github.com/UncleAndy/green-api-test/internal/templates"
)

type Server struct {
	cfg *config.ServerConfig

	templates map[string]*template.Template

	client *greenapi.Client
}

func New(cfg *config.ServerConfig, greenClient *greenapi.Client) (*Server, error) {
	server := &Server{
		cfg:    cfg,
		client: greenClient,
	}

	server.prepareTemplates()

	return server, nil
}

func (s *Server) Run() {
	httpRouter := router.New()

	httpRouter.GET("/", s.Index)

	httpRouter.GET("/get_settings", s.GetSettings)
	httpRouter.GET("/get_state_instance", s.GetStateInstance)

	server := &fasthttp.Server{
		Handler:            httpRouter.Handler,
		ReadBufferSize:     4096,
		MaxRequestBodySize: 1024 * 1024 * 50,
	}

	log.Printf("server started at %v:%v", s.cfg.Host, s.cfg.Port)
	err := server.ListenAndServe(fmt.Sprintf("%v:%v", s.cfg.Host, s.cfg.Port))
	if err != nil {
		log.Fatalf("unexpected error in server: %v", err)
	}
}

func (s *Server) prepareTemplates() {
	mainTmpl := `{{define "main" }} {{ template "base" . }} {{ end }}`
	templatesList := make(map[string]*template.Template)

	mainTemplate, err := template.New("main").Parse(mainTmpl)
	if err != nil {
		log.Fatal("failed to parse main template", err)
	}

	contentFiles, err := templates.FileNames(templates.ContentFS)
	if err != nil {
		log.Fatal("failed to get file names for blocks", err)
	}

	contentFilesContent := make([]string, len(contentFiles))
	for i := 0; i < len(contentFiles); i++ {
		data, err := templates.ContentFS.ReadFile(contentFiles[i])
		if err != nil {
			log.Fatal("failed to read template file", err)
		}
		contentFilesContent[i] = string(data)
	}

	incFiles, err := templates.FileNames(templates.BaseFS)
	if err != nil {
		log.Fatal("failed to get file names for templates", err)
	}
	incFilesContent := make([]string, len(incFiles))
	for i := 0; i < len(incFiles); i++ {
		data, err := templates.BaseFS.ReadFile(incFiles[i])
		if err != nil {
			log.Panic("failed to read template file", err)
		}
		incFilesContent[i] = string(data)
	}

	for i, file := range contentFiles {
		fileName := filepath.Base(file)
		files := append(incFilesContent, contentFilesContent[i]) // nolint
		templatesList[fileName], err = mainTemplate.Clone()
		if err != nil {
			log.Fatal("got error while trying to clone mainTemplate", err)
		}

		templatesList[fileName] = template.Must(templatesList[fileName].Parse(strings.Join(files, "\n")))

		log.Printf("template loaded added %s", fileName)
	}

	s.templates = templatesList

	log.Print("templates loaded successful")
}

func (s *Server) renderTemplate(ctx *fasthttp.RequestCtx, name string, data interface{}) {
	tmpl, ok := s.templates[name]
	if !ok {
		log.Printf("template not found %s", name)

		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	err := tmpl.Execute(ctx, data)
	if err != nil {
		log.Printf("template execution failed: %s - %v", name, err)

		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}
}
