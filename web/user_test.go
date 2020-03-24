package web

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/labstack/echo"
	"github.com/lualfe/ecommerce/app/inmemory"
	"github.com/lualfe/ecommerce/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var (
	user1 = &models.User{
		ID:        1,
		Email:     "test@example.com",
		Password:  "123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	user2 = &models.User{
		ID:        2,
		Email:     "test2@example.com",
		Password:  "456",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
)

type UserSuite struct {
	suite.Suite

	e *echo.Echo
	w *Web
}

func (s *UserSuite) SetupSuite() {
	in := inmemory.InMemory{}
	in[1] = user1
	in[2] = user2
	s.w = NewWeb(in, nil)
	s.e = echo.New()
}

func TestUserInit(t *testing.T) {
	suite.Run(t, new(UserSuite))
}

func (s *UserSuite) TestInsertUser() {
	req := httptest.NewRequest(http.MethodPost, "/register", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)
	c.Set("user_id", 1)
	require.Error(s.T(), s.w.InsertUser(c), "user cannot surpass if a session is active")

	form := url.Values{}
	form.Add("email", "test@example.com")
	req = httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	req.Form = form
	rec = httptest.NewRecorder()
	c = s.e.NewContext(req, rec)
	require.Error(s.T(), s.w.InsertUser(c), "user already exists and should't be able to insert again")

	form = url.Values{}
	form.Add("email", "test3@example.com")
	form.Add("password", "789")
	req = httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(form.Encode()))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	req.Form = form
	rec = httptest.NewRecorder()
	c = s.e.NewContext(req, rec)
	if assert.NoError(s.T(), s.w.InsertUser(c)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
	}
}
