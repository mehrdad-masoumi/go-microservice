package node_repo

import (
	"database/sql"
	"mlm/entity"
	"mlm/pkg/error_msg"
	"mlm/pkg/richerror"
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
	const op = "node_repo.Create"

	res, err := d.db.Exec(`insert into nodes(id, parent_id, line, lft_referral, rgt_referral, ancestry) values (?, ?, ?, ?, ?, ?)`,
		node.ID,
		node.ParentId,
		node.Line,
		node.LftReferral,
		node.RgtReferral,
		node.Ancestry,
	)

	if err != nil {
		return entity.Node{}, richerror.New(op).
			WithErr(err).
			WithMeta(map[string]interface{}{"data": node})
	}

	id, _ := res.LastInsertId()

	node.ID = uint(id)

	return node, nil
}

func (d DB) FindNodeByReferral(referral string) (entity.Node, error) {
	const op = "node_repo.FindNodeByReferral"

	var node entity.Node

	row := d.db.QueryRow(`select id, ancestry from nodes where nodes.lft_referral = ? or nodes.rgt_referral= ?`,
		referral,
		referral,
	)
	err := row.Scan(&node.ID, &node.Ancestry)
	if err != nil {
		if err == sql.ErrNoRows {
			return entity.Node{}, richerror.New(op).WithErr(err).
				WithMessage(error_msg.ErrorMsgNotFound).WithKind(richerror.KindNotFound)
		}

		return entity.Node{}, richerror.New(op).WithErr(err).
			WithMessage(error_msg.ErrorMsgCantScanQueryResult).WithKind(richerror.KindUnexpected)
	}

	return node, nil
}

func (d DB) Delete(id uint) (bool, error) {

	const op = "node_repo.IsEmailUnique"

	_, err := d.db.Exec(`delete from nodes where id = ?`, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New(op).WithErr(err).WithMessage(error_msg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return true, nil

}
