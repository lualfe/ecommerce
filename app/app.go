package app

import (
	"github.com/lualfe/ecommerce/models"
)

// Repository ...
type Repository interface {
	InsertUser(user *models.User) (*models.User, error)
	GetUserByEmail(user *models.User) (*models.User, error)
}

// Migration ...
type Migration interface {
	Migrate()
}
