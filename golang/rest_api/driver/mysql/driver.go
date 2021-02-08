package driver

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type DB struct {
	SQL *sql.DB
}

var dbConn = &DB{}

func ConnectSQL(host, port, uname, pass, dbname string) (*DB, error) {
	DBSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		uname, pass, host, port, dbname)
	db, err := sql.Open("mysql", DBSource)
	if err != nil {
		log.Panicln(err)
	}
	dbConn.SQL = db
	return dbConn, err
}
