package db

import (
	"github.com/adityaladwa/todoapp/db"
	"github.com/go-pg/pg/v10"
)

func Connect() *pg.DB {
	DbCofig := db.GetDbConfig()
	return pg.Connect(&pg.Options{
		User:     DbCofig.Username,
		Database: DbCofig.DbName,
	})
}
