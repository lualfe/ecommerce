package pg

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/lualfe/ecommerce/app"
	"github.com/lualfe/ecommerce/models"
)

// InsertUser inerts a new user into the database
func (a *Postgres) InsertUser(user *models.User) (*models.User, error) {
	if err := a.Instance.Create(&user).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return user, nil
}

// GetUserByEmail gets an user by its email
func (a *Postgres) GetUserByEmail(user *models.User) (*models.User, error) {
	u := &models.User{}
	if err := a.Instance.Where("email = ?", user.Email).Find(&u).Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return u, nil
}

// CreateUserRepository initializes an user respository for testing
func CreateUserRepository(db *gorm.DB) app.Repository {
	return &Postgres{db}
}
