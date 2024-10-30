package user_repo

import (
	"database/sql"
	"mlm/entity"
	"mlm/pkg/error_msg"
	"mlm/pkg/richerror"
)

type DB struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *DB {
	return &DB{
		db: db,
	}
}

func (d DB) Create(user entity.User) (entity.User, error) {

	const op = "user_repo.Create"
	res, err := d.db.Exec(`insert into users(email, phone_number, role, password) values (?, ?, ?, ?)`,
		user.Email,
		user.PhoneNumber,
		user.Role,
		user.Password,
	)

	if err != nil {
		return entity.User{}, richerror.New(op).
			WithErr(err).
			WithMeta(map[string]interface{}{"data": user})
	}

	id, err := res.LastInsertId()
	if err != nil {
		return entity.User{}, richerror.New(op).WithErr(err).WithMeta(map[string]interface{}{"data": user})
	}

	user.ID = uint(id)

	return user, nil
}

func (d DB) IsEmailUnique(email string) (bool, error) {
	const op = "node_repo.IsEmailUnique"

	row := d.db.QueryRow(`select * from users where email = ?`, email)

	err := row.Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			return true, nil
		}

		return false, richerror.New(op).WithErr(err).WithMessage(error_msg.ErrorMsgCantScanQueryResult).
			WithKind(richerror.KindUnexpected)
	}

	return false, nil

}
