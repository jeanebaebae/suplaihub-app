package api

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	port string
}

func NewServer(port string) *Server{
	return &Server{
		port: port,
	}
}

func (s *Server) MountRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","message":"Suplaihub API Berjalan!"}`))
	})

	return r
}

func (s *Server) Run() error {
	handler := s.MountRoutes()
	fmt.Printf("Server running on port :%s\n", s.port)
	return http.ListenAndServe(":"+s.port, handler)
}