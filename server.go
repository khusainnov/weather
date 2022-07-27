package weather

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    time.Second * 15,
		WriteTimeout:   time.Second * 15,
		MaxHeaderBytes: 1 << 20, // 1 MB
	}

	// TODO: mux := http.NewServeMux()

	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
