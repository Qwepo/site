package db

import (
	"errors"
	"time"
)

type dbSessions interface {
	SessionCreate(*SessionFull) error

	SessionDeleteByID(ctx Context, id PK) error
	SessionsDeleteOthers(ctx Context, userID, currentSessionID PK) error

	SessionUpdateByID(ctx Context, id PK) error

	SessionActiveByToken(ctx Context, token string) (*SessionFull, error)
	SessionsByUserID(ctx Context, userID PK, expiredLimit int) ([]*Session, error)
}

var (
	ErrSessionTokenNoSet = errors.New("session token no set")
	ErrSessionUserNoSet  = errors.New("session user no set")
)

type Session struct {
	ID        int64 `reindex:"id"              json:"id"`
	CreatedAt int64 `reindex:"created_at,-"    json:"createdAt"`
	UpdatedAt int64 `reindex:"updated_at,tree" json:"updatedAt"`
	DeletedAt int64 `reindex:"deleted_at"      json:"deletedAt"`

	UserID int64  `reindex:"user_id"       json:"userID,omitempty"`
	IP     string `reindex:"ip,-"          json:"ip"`
	OS     string `reindex:"os,-"          json:"os"`
	Mobile bool   `reindex:"mobile,-"      json:"mobile"`
}

type SessionFull struct {
	Session
	Token string `reindex:"token,,pk" json:"token,omitempty"`
}

var sessionFields = []string{
	"id", "createdAt", "updatedAt", "deletedAt", "userID",
	"ip", "os", "mobile", "browser", "browserVer",
}

func (db *client) SessionCreate(s *SessionFull) error {
	switch {
	case s.Token == "":
		return ErrSessionTokenNoSet
	case s.UserID == 0:
		return ErrSessionUserNoSet
	}
	s.CreatedAt = time.Now().Unix()
	return db.Create(nsSession, s)
}

func (db *client) SessionDeleteByID(ctx Context, id PK) error {
	q := db.Query(nsSession).
		WhereInt64("id", EQ, id).
		Set("deletedAt", nowUnix())
	return db.Updates(ctx, q)
}

func (db *client) SessionsDeleteOthers(ctx Context, userID, currentSessionID PK) error {
	q := db.Query(nsSession).
		WhereInt64("userID", EQ, userID).
		WhereInt64("deletedAt", EQ, 0).
		Not().WhereInt64("id", EQ, currentSessionID).
		Set("deletedAt", nowUnix())
	return db.Updates(ctx, q)
}

func (db *client) SessionUpdateByID(ctx Context, id PK) error {
	q := db.Query(nsSession).
		WhereInt64("id", EQ, id).
		Set("updatedAt", nowUnix())
	return db.Updates(ctx, q)
}

func (db *client) SessionActiveByToken(ctx Context, token string) (*SessionFull, error) {
	q := db.Query(nsSession).
		WhereInt64("deletedAt", EQ, 0).
		WhereString("token", EQ, token)
	return db.sessionFullGet(ctx, q)
}

func (db *client) SessionsByUserID(ctx Context, userID PK, expiredLimit int) ([]*Session, error) {
	active, err := db.Query(nsSession).
		Select(sessionFields...).
		WhereInt64("userID", EQ, userID).
		WhereInt64("deletedAt", EQ, 0).
		Sort("updatedAt", true).
		ExecCtx(ctx).FetchAll()
	if err != nil {
		return nil, err
	}
	var expired []interface{}
	if delta := expiredLimit - len(active); delta > 0 {
		expired, _ = db.Query(nsSession).
			Select(sessionFields...).
			WhereInt64("userID", EQ, userID).
			Not().WhereInt64("deletedAt", EQ, 0).
			Sort("updatedAt", true).Limit(delta).
			ExecCtx(ctx).FetchAll()
	}
	sessions := make([]*Session, len(active)+len(expired))
	for i := range active {
		sessions[i] = &active[i].(*SessionFull).Session
	}
	for i := range expired {
		sessions[len(active)+i] = &expired[i].(*SessionFull).Session
	}
	return sessions, nil
}

// =============================================================

func (db *client) sessionFullGet(ctx Context, q Query) (*SessionFull, error) {
	item, err := db.Get(ctx, q)
	if item == nil {
		return nil, err
	}
	return item.(*SessionFull), nil
}
