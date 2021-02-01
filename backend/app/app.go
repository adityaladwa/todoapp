package app

import (
	db "github.com/adityaladwa/todoapp/db/connect"
	"github.com/go-pg/pg/v10"
)

type App struct {
	DB *pg.DB
}

func (a *App) Init() error {
	conn := db.Connect()
	a.DB = conn
	return nil
}
