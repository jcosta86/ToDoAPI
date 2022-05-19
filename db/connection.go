package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.jcosta86.com/todoapi/configs"
)

// OpenConnection opens a connection to the database and returns it
func OpenConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DataBase)

	con, err := sql.Open("postgres", strConn)
	if err != nil {
		return nil, err
	}
	err = con.Ping()
	return con, err
}
