package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/lualfe/ecommerce/app/inmemory"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

var jsonResponse = `{"status":"ok"}`

type PurchaseSuite struct {
	suite.Suite

	e *echo.Echo
	w *Web
}

func (s *PurchaseSuite) SetupSuite() {
	in := inmemory.InMemory{}
	s.w = NewWeb(in, nil)
	s.e = echo.New()
	viper.SetConfigName("default")    // name of config file (without extension)
	viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("../config/") // path to look for the config file in
	err := viper.ReadInConfig()       // Find and read the config file
	if err != nil {                   // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}
}

func TestInit(t *testing.T) {
	suite.Run(t, new(PurchaseSuite))
}

func (s *PurchaseSuite) TestNewPurchase() {
	form := url.Values{}
	form.Add("zip", "04049060")
	form.Add("billing_zip", "04421100")
	req := httptest.NewRequest(http.MethodPost, "/purchase", strings.NewReader(form.Encode()))
	req.Form = form
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec := httptest.NewRecorder()
	c := s.e.NewContext(req, rec)
	require.Error(s.T(), s.w.NewPurchase(c), "distance between both addresses is greater than 5000 meters, so an error is expected")

	form = url.Values{}
	form.Add("zip", "04049060")
	form.Add("billing_zip", "04049050")
	req = httptest.NewRequest(http.MethodPost, "/purchase", strings.NewReader(form.Encode()))
	req.Form = form
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
	rec = httptest.NewRecorder()
	c = s.e.NewContext(req, rec)
	if assert.NoError(s.T(), s.w.NewPurchase(c)) {
		assert.Equal(s.T(), http.StatusOK, rec.Code)
		assert.True(s.T(), strings.Contains(rec.Body.String(), `"status":"ok"`))
	}
}
