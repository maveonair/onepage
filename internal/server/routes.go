package server

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//go:embed templates/*
var templates embed.FS

//go:embed resources/*
var resources embed.FS

func (s *server) routes() error {
	s.router.HandleFunc("/_health", s.handleHealth()).Methods(http.MethodGet)

	s.router.HandleFunc("/", s.handleIndex()).Methods(http.MethodGet)
	s.router.HandleFunc("/edit", s.handleEdit()).Methods(http.MethodGet)
	s.router.HandleFunc("/update", s.handleUpdate()).Methods(http.MethodPost)

	resourcesFs, err := fs.Sub(resources, "resources")
	if err != nil {
		return err
	}

	s.router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.FS(resourcesFs))))

	return nil
}

func (s *server) handleHealth() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	}
}

func (s *server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(templates, "templates/index.html", "templates/base.html")
		if err != nil {
			log.WithError(err).Error("failed to parse template")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		content, err := s.MarkdownToHTML()
		if err != nil {
			log.WithError(err).Error("failed to render content")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := tmpl.ExecuteTemplate(w, "base", template.HTML(content)); err != nil {
			log.WithError(err).Error("failed to render template")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) handleEdit() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFS(templates, "templates/edit.html", "templates/base.html")
		if err != nil {
			log.WithError(err).Error("failed to parse template")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		content, err := s.ReadPageFile()
		if err != nil {
			log.WithError(err).Error("failed to render content")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := tmpl.ExecuteTemplate(w, "base", string(content)); err != nil {
			log.WithError(err).Error("failed to render template")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func (s *server) handleUpdate() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.WithError(err).Error("failed to parse form")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		content := r.FormValue("content")

		if err := s.WritePageFile([]byte(content)); err != nil {
			log.WithError(err).Error("failed to write file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusFound)
	}
}
