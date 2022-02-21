package db

import (
	"context"
	"moment/pkg/utilty"
	"time"

	"gitlab.com/knopkalab/go/database/reindexer"
	"gitlab.com/knopkalab/go/logger"
)

const (
	nsUser    = "users"
	nsSession = "sessions"
)

type (
	Context = context.Context
	Query   = reindexer.Query

	PK   = utilty.PK
	Unix = utilty.Unix
)

var namespaces = map[string]interface{}{
	nsUser:    new(User),
	nsSession: new(SessionFull),
}

type DB interface {
	dbUsers
	dbSessions
}

type Client interface {
	DB
	Close() error
	Migrate() error
}

const (
	EQ     = reindexer.EQ     // =
	GT     = reindexer.GT     // >
	LT     = reindexer.LT     // <
	GE     = reindexer.GE     // >= (GT|EQ)
	LE     = reindexer.LE     // <= (LT|EQ)
	SET    = reindexer.SET    // one of set in []
	ALLSET = reindexer.ALLSET // all of set in []
	RANGE  = reindexer.RANGE  // value in RANGE from,to
	ANY    = reindexer.ANY    // any value
	EMPTY  = reindexer.EMPTY  // empty value or len(value array) == 0
	LIKE   = reindexer.LIKE   // string like pattern
)

var (
	ErrPrimaryKey = reindexer.ErrPrimaryKey
	ErrNoAffected = reindexer.ErrNoAffected
)

type client struct {
	*reindexer.Conn
}

func Open(conf reindexer.Config, log logger.Logger) (Client, error) {
	conn, err := reindexer.Open(conf, log)
	return &client{conn}, err
}

func (db *client) Migrate() error {
	return db.Conn.Migrate(namespaces, false)
}
func nowUnix() int64 {
	return time.Now().Unix()
}
