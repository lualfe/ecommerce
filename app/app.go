package app

import (
	"github.com/lualfe/ecommerce/app/pg"
	"github.com/spf13/viper"
)

// App model
type App struct {
	PG *pg.Postgres
}

// NewApp returns a new instance of App object
func NewApp() *App {
	a := &App{}
	a = a.NewPostgres(viper.GetString("PG_CONNECT"))
	return a
}
