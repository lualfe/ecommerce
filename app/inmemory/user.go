package inmemory

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/lualfe/ecommerce/models"
)

// InsertUser ...
func (a InMemory) InsertUser(user *models.User) (*models.User, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	i := 1
	for {
		_, ok := a[i]
		if ok {
			break
		}
		i++
	}
	for _, v := range a {
		if v.Email == user.Email {
			return nil, echo.NewHTTPError(http.StatusBadRequest, "user already exists")
		}
	}
	user.ID = i
	a[i] = user
	return a[i], nil
}

// GetUserByEmail ...
func (a InMemory) GetUserByEmail(user *models.User) (*models.User, error) {
	for _, v := range a {
		if v.Email == user.Email {
			return v, nil
		}
	}
	return nil, echo.NewHTTPError(http.StatusBadRequest, "record not found")
}
