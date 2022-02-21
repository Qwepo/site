package web

import (
	"moment/pkg/db"

	"gitlab.com/knopkalab/go/http"
)

// Request is web.context wrapper
type Request struct {
	*http.Request

	Session *db.SessionFull
	User    *db.User
}

type jsonFormChecker interface {
	Check() bool
}

// ParseJSONForm and call Check on success
func (r *Request) ParseJSONForm(dst jsonFormChecker) bool {
	return r.ParseJSON(dst) == nil && dst.Check()
}
