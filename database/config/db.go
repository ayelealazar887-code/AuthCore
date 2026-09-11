package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	db *sql.DB
}

func NewConnection(db *sql.DB) *Mysql {
	return &Mysql{
		db : db,
	}
}
func DbConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:rootadmin@tcp(127.0.0.1:3306)/authGO?parseTime=true")
	if err !=nil {
		panic("failed to connect: "+ err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("fialed to ping database: " + err.Error())
	}

	return db
}
