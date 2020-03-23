package app

import (
	"github.com/jinzhu/gorm"
	"github.com/lualfe/ecommerce/app/pg"

	//needed for postgres initialization
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// Necessary in order to check for transaction retry error codes.
	_ "github.com/lib/pq"
)

// NewPostgres initializes a postgres instance
func (a *App) NewPostgres(conn string) *App {
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	pg := &pg.Postgres{
		Instance: db,
	}
	a.PG = pg
	return a
}
