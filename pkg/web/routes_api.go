package web

import (
	"moment/pkg"
	"moment/pkg/db"
	phoneapi "moment/pkg/phoneAPI"

	"gitlab.com/knopkalab/go/http"
	"gitlab.com/knopkalab/go/logger"
)

func registerAPI(r *http.Router, log logger.Logger, ctx db.Context, ctr *Controllers, ucl phoneapi.Ucaller, conf *pkg.Config) {
	r.Post("/account/auth", AuthView(ctr.User, conf.Ucaller, log))
	r.Post("/account/login", LoginView(ctr.User, ctr.Session, log, conf))
	r.AddMiddleware(authRequiredMiddleware)

	r.Get("/account/test", Tesst)

}
