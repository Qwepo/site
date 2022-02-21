package web

import (
	"context"
	"moment/pkg"

	"gitlab.com/knopkalab/go/http"
	"gitlab.com/knopkalab/go/logger"
)

type Server struct {
	Server *http.Server
}

func (s *Server) Stop() error {
	return s.Server.Close()
}

func StartServer(conf *pkg.Config, log logger.Logger, ctr *Controllers) (*Server, error) {
	r := newRouter(log)
	if err := registerRoutes(r, log, context.TODO(), ctr, conf); err != nil {
		return nil, err
	}

	httpServer, err := http.StartServer(conf.Server, log, r)
	if err != nil {
		return nil, err
	}

	handler := &Server{
		Server: httpServer,
	}
	return handler, nil
}
