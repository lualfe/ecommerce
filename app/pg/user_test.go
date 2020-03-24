package pg

import (
	"database/sql"
	"reflect"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/lualfe/ecommerce/app"
	"github.com/lualfe/ecommerce/models"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type UserSuite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo app.Repository
	user *models.User
}

func (s *UserSuite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("postgres", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repo = CreateUserRepository(s.DB)
}

func TestInit(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (s *UserSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func (s *UserSuite) TestGetUserByEmail() {
	var (
		id        = 1
		email     = "test@example.com"
		password  = "123"
		createdAt = time.Now()
		updatedAt = time.Now()
	)
	user := &models.User{
		ID:        id,
		Email:     email,
		Password:  password,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(`
	SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL AND ((email = $1))
	`)).
		WithArgs(user.Email).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "password", "created_at", "updated_at", "deleted_at"}).AddRow(user.ID, user.Email, user.Password, user.CreatedAt, user.UpdatedAt, user.DeletedAt))

	userAfter, err := s.repo.GetUserByEmail(user)
	require.NoError(s.T(), err)
	require.True(s.T(), reflect.DeepEqual(user, userAfter), "user returning from GetUserByEmail must match mock")
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

	s.mock.ExpectBegin()
	s.mock.ExpectQuery(regexp.QuoteMeta(`
	INSERT INTO "users" ("email","password","created_at","updated_at","deleted_at") 
	VALUES ($1,$2,$3,$4,$5) RETURNING "users"."id"
	`)).
		WithArgs(user.Email, user.Password, user.CreatedAt, user.UpdatedAt, sql.NullTime{}).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(user.ID))
	s.mock.ExpectCommit()

	_, err := s.repo.InsertUser(user)
	require.NoError(s.T(), err)
}
