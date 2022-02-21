package web

import (
	"gitlab.com/knopkalab/go/http"
	"gitlab.com/knopkalab/go/logger"
)

// View is route request view
type View func(*Request) error

// Viewer is interface returns view
type Viewer interface {
	View() View
}

func newRouter(log logger.Logger) *http.Router {
	router := http.NewRouter(log)
	router.ViewConverter = func(view interface{}) http.View {
		switch view := view.(type) {
		case func(*Request) error:
			return func(r *http.Request) error {
				return view(r.Payload.(*Request))
			}
		case View:
			return func(r *http.Request) error {
				return view(r.Payload.(*Request))
			}
		case Viewer:
			v := view.View()
			return func(r *http.Request) error {
				return v(r.Payload.(*Request))
			}
		default:
			return http.DefaultViewConverter(view)
		}
	}
	return router
}
