package web

import (
	"moment/pkg"
	"moment/pkg/services"

	"gitlab.com/knopkalab/go/http"
)

func sessionMiddleware(conf *pkg.Config, sessions services.Sessions, users services.User) http.Middleware {
	return func(next http.View) http.View {
		return func(req *http.Request) error {
			// get session cookie
			cookie, err := req.Cookie(conf.Sessions.CookieName)
			if err != nil {
				return next(req)
			}

			r := req.Payload.(*Request)

			r.Session, err = sessions.GetActiveByToken(r, cookie.Value)

			if err != nil {
				r.Log.Err(err).Msg("cannot fetch user session")
			}

			if r.Session == nil {
				return next(req)
			}

			r.User, err = users.GetUserByID(r.Session.UserID, r)
			if err != nil {
				r.Log.Err(err).Msg("cannot fetch user")
			}

			if r.User != nil {
				err = sessions.UpdateByID(r, r.Session.ID)
				if err != nil {
					r.Log.Err(err).Msg("cannot update session")
				}

				// update session cookie
				cookie.Path = "/"
				cookie.HttpOnly = true
				cookie.MaxAge = conf.Sessions.CookieMaxAge.Seconds()
				r.Response.SetCookie(cookie)

				r.Log = r.Log.With().
					Int64("user_id", r.User.ID).
					Int64("session_id", r.Session.ID).Logger()
			} else {
				r.Log = r.Log.With().
					Int64("session_id", r.Session.ID).Logger()
			}

			return next(req)
		}
	}
}
