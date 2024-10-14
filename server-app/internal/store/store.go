package store

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Storage struct {
	db     *sql.DB
	Rep    *Rep
	logger *logrus.Logger
}

func New(db *sql.DB, l *logrus.Logger) *Storage {
	return &Storage{
		db:     db,
		logger: l,
	}
}

func (st *Storage) Repository() Rep {
	if st.Rep != nil {
		return *st.Rep
	}

	st.Rep = &Rep{
		storage: st,
	}

	return *st.Rep
}
