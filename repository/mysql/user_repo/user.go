package user_repo

import (
	"database/sql"
	"mlm/entity"
)

type DB struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

func (d DB) Create(node entity.User) (entity.User, error) {
	//TODO implement me
	panic("implement me")
}
