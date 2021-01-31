package db

import "github.com/go-pg/pg/v10"

func Connect() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     "adityaladwa",
		Database: "todoapp",
	})
}
