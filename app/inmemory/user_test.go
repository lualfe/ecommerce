package inmemory

import (
	"testing"
	"time"

	"github.com/lualfe/ecommerce/app"
	"github.com/lualfe/ecommerce/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserSuite struct {
	suite.Suite
	inmemory InMemory

	repo app.Repository
	user *models.User
}

func (s *UserSuite) SetupSuite() {
	s.inmemory = make(map[int]*models.User)
	s.inmemory[1] = &models.User{
		Email:     "test1@example.com",
		Password:  "123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.inmemory[2] = &models.User{
		Email:     "test2@example.com",
		Password:  "456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	s.inmemory[3] = &models.User{
		Email:     "test3@example.com",
		Password:  "789",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func TestInit(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (s *UserSuite) TestGetUserByEmail() {
	var (
		email = "test1@example.com"
	)
	user := &models.User{
		Email: email,
	}

	_, err := s.inmemory.GetUserByEmail(user)
	require.NoError(s.T(), err)
}

func (s *UserSuite) TestInsertUser() {
	var (
		email    = "test@example.com"
		password = "123"
	)
	user := &models.User{
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := s.inmemory.InsertUser(user)
	require.NoError(s.T(), err)
}
