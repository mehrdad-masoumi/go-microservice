package node_repo

import (
	"database/sql"
	"mlm/entity"
)

type DB struct {
	db *sql.DB
}

func NewNodeRepository(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

func (d DB) Create(node entity.Node) (entity.Node, error) {
	//TODO implement me
	panic("implement me")
}
func (d DB) FindNodeByReferral(referral string) (entity.Node, error) {
	//TODO implement me
	panic("implement me")
}
