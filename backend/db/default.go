package db

import (
	"errors"
	"strconv"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	MySQL *sqlx.DB
}

func Open(driverName string, username string, password string, dbname string, host string, port int) (*sqlx.DB, error) {

	switch driverName {
	case "mysql":
		return sqlx.Connect("mysql", username+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+dbname+"?")
	}

	return nil, errors.New("")

}
