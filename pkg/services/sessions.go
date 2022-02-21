package services

import (
	"moment/pkg/db"

	"gitlab.com/knopkalab/go/logger"
	"gitlab.com/knopkalab/go/utils"
)

type Sessions interface {
	Create(*SessionCreateRequest) (*db.SessionFull, error)

	DeleteByID(ctx Context, id PK) error
	DeleteOthers(ctx Context, userID, currentSessionID PK) error

	UpdateByID(ctx Context, id PK) error

	GetActiveByToken(ctx Context, token string) (*db.SessionFull, error)
	GetListByUserID(ctx Context, userID PK, expiredLimit int) ([]*db.Session, error)
}

type SessionServices struct {
	log logger.Logger
	db  db.DB
}

func (s *SessionServices) Create(req *SessionCreateRequest) (session *db.SessionFull, err error) {
	session = req.toFullSession()
	for err = db.ErrPrimaryKey; err == db.ErrPrimaryKey; {
		session.Token = utils.RandString(64)
		err = s.db.SessionCreate(session)
	}
	return session, err
}

func (s *SessionServices) DeleteByID(ctx Context, id PK) error {
	return s.db.SessionDeleteByID(ctx, id)
}

func (s *SessionServices) DeleteOthers(ctx Context, userID, currentSessionID PK) error {
	return s.db.SessionsDeleteOthers(ctx, userID, currentSessionID)
}

func (s *SessionServices) UpdateByID(ctx Context, id PK) error {
	return s.db.SessionUpdateByID(ctx, id)
}

func (s *SessionServices) GetActiveByToken(ctx Context, token string) (*db.SessionFull, error) {
	return s.db.SessionActiveByToken(ctx, token)
}

func (s *SessionServices) GetListByUserID(ctx Context, userID PK, expiredLimit int) ([]*db.Session, error) {
	return s.db.SessionsByUserID(ctx, userID, expiredLimit)
}

func NewSessions(log logger.Logger, db db.DB) Sessions {
	return &SessionServices{log: log, db: db}
}
