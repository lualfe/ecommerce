package web

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/lualfe/ecommerce/models"
	"github.com/lualfe/ecommerce/utils"
	"github.com/spf13/viper"
)

// InsertUser inserts an new user into the database.
func (w *Web) InsertUser(c echo.Context) error {
	if c.Get("user_id") != nil {
		return echo.NewHTTPError(http.StatusPermanentRedirect, "user already logged in")
	}

	user := &models.User{}
	c.Bind(user)
	user, err := w.App.PG.InsertUser(user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	claims := &utils.JWTClaims{
		UserID: user.ID,
	}
	claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(viper.GetString("jwt_key")))
	if err != nil {
		return err
	}
	c.Set("user_id", user.ID)
	type finalResponse struct {
		User  *models.User `json:"user"`
		Token string       `json:"access_token"`
	}
	response := &finalResponse{
		User:  user,
		Token: t,
	}
	c.JSON(http.StatusOK, response)
	return nil
}

// LoginUser logs an user in
func (w *Web) LoginUser(c echo.Context) error {
	if c.Get("user_id") != nil {
		return echo.NewHTTPError(http.StatusPermanentRedirect, "user already logged in")
	}

	user := &models.User{}
	c.Bind(user)
	user, err := w.App.PG.LoginUser(user)
	if err != nil {
		return err
	}
	claims := &utils.JWTClaims{
		UserID: user.ID,
	}
	claims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(viper.GetString("jwt_key")))
	if err != nil {
		return err
	}
	c.Set("user_id", user.ID)
	type finalResponse struct {
		User  *models.User `json:"user"`
		Token string       `json:"access_token"`
	}
	response := &finalResponse{
		User:  user,
		Token: t,
	}
	c.JSON(http.StatusOK, response)
	return nil
}
