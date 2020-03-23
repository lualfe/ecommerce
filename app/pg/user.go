package pg

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lualfe/ecommerce/models"
	"github.com/lualfe/ecommerce/utils"
)

func (a *Postgres) InsertUser(user *models.User) (*models.User, error) {
	hash, err := utils.HashPassword(user.Password)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "could not register user")
	}
	user.Password = hash
	if err := a.Instance.Create(&user).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return user, nil
}

func (a *Postgres) LoginUser(user *models.User) (*models.User, error) {
	u := &models.User{}
	if err := a.Instance.Where("email = ?", user.Email).Find(&u).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if !utils.ComparePassword(user.Password, u.Password) {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "email or password incorrect")
	}
	return u, nil
}
