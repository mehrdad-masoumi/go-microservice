package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"mlm/config"
)

func Connect(config config.Mysql) (*sql.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Username, config.Password, config.Host, config.Port, config.DB)

	conn, err := sql.Open("mysql", dsn)

	err = conn.Ping()

	return conn, err

}

func Close(conn *sql.DB) error {
	return conn.Close()
}
