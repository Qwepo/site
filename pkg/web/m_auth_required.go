package web

import "gitlab.com/knopkalab/go/http"

func authRequiredMiddleware(next http.View) http.View {
	return func(req *http.Request) error {
		r := req.Payload.(*Request)
		if r.User == nil {
			return r.Response.StatusUnauthorized()
		}
		return next(req)
	}
}
