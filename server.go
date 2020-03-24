package main

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lualfe/ecommerce/app/pg"
	"github.com/lualfe/ecommerce/web"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("default")   // name of config file (without extension)
	viper.SetConfigType("yaml")      // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./config/") // path to look for the config file in
	err := viper.ReadInConfig()      // Find and read the config file
	if err != nil {                  // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func main() {
	e := echo.New()
	postgres := pg.NewPostgres(viper.GetString("PG_CONNECT"))
	w := web.NewWeb(postgres, postgres)
	w.MG.Migrate()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "<-- Method: ${method} | URI: ${uri} | Status: ${status} | Latency: ${latency_human} -->\n",
	}))
	w.InitRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
