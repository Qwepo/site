package web

import (
	"sync"

	"gitlab.com/knopkalab/go/http"
)

type requestMiddleware struct {
	pool *sync.Pool
}

func newRequestMiddleware() *requestMiddleware {
	return &requestMiddleware{
		pool: &sync.Pool{New: func() interface{} { return new(Request) }},
	}
}

func (m *requestMiddleware) getRequest(origin *http.Request) *Request {
	r := m.pool.Get().(*Request)

	r.Request = origin

	r.Session = nil
	r.User = nil

	return r
}

func (m *requestMiddleware) View(next http.View) http.View {
	return func(req *http.Request) error {
		r := m.getRequest(req)
		defer m.pool.Put(r)

		req.Payload = r

		return next(req)
	}
}
