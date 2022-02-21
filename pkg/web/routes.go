package web

import (
	"moment/pkg"
	"moment/pkg/db"

	"gitlab.com/knopkalab/go/http"
	"gitlab.com/knopkalab/go/http/middlewares"
	"gitlab.com/knopkalab/go/logger"
)

func registerRoutes(r *http.Router, log logger.Logger, ctx db.Context, ctr *Controllers, conf *pkg.Config) error {
	r.AddMiddleware(middlewares.Logger)
	r.AddMiddleware(
		&middlewares.CSRF{},
		newRequestMiddleware(),
		sessionMiddleware(conf, ctr.Session, ctr.User),
	)
	registerAPI(r.Fork("/api"), log, ctx, ctr, conf.Ucaller, conf)
	return nil
}
