package db

import (
	"github.com/adityaladwa/todoapp/db"
	"github.com/go-pg/pg/v10"
)

var DBConn *pg.DB

func Connect() {
	DbCofig := db.GetDbConfig()
	DBConn = pg.Connect(&pg.Options{
		User:     DbCofig.Username,
		Database: DbCofig.DbName,
	})
}
