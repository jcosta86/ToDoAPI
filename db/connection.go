package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	c "github.jcosta86.com/todoapi/configs"
)

func OpenConnection() (*sql.DB, error) {
	conf := c.GetDB()

	strConn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DataBase)

	con, err := sql.Open("postgres", strConn)
	if err != nil {
		return nil, err
	}
	err = con.Ping()
	return con, err
}
