package web

import (
	"moment/pkg/db"
	"moment/pkg/services"

	"gitlab.com/knopkalab/go/logger"
)

type Controllers struct {
	Session services.Sessions
	User    services.User
}

func NewController(log logger.Logger, db db.DB) *Controllers {
	session := services.NewSessions(log, db)
	users := services.NewUser(log, db)

	return &Controllers{
		Session: session,
		User:    users,
	}

}
