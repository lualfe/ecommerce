package web

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/lualfe/ecommerce/utils"
	"github.com/spf13/viper"
)

// InitRoutes initializes all application routes
func (w *Web) InitRoutes(e *echo.Echo) {
	e.POST("/register", w.InsertUser, utils.CheckToken)
	e.POST("/login", w.LoginUser, utils.CheckToken)

	e.POST("/purchase", w.NewPurchase, middleware.JWT([]byte(viper.GetString("jwt_key"))))
}
