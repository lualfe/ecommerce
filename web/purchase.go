package web

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/lualfe/ecommerce/models"
	"github.com/lualfe/ecommerce/utils"
	"github.com/spf13/viper"
)

// NewPurchase checks if a new purchase is valid. If so, it will be confirmed.
func (w *Web) NewPurchase(c echo.Context) error {
	address := &models.Address{}
	billingAddress := &models.BillingAddress{}
	c.Bind(address)
	c.Bind(billingAddress)
	if address.ZIP == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "ZIP information is required for the address")
	}
	if billingAddress.BillingZIP == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "ZIP information is required for the billing address")
	}

	// current location info coming from HTML5 geolocation api
	currentLat, err := strconv.ParseFloat(c.QueryParam("lat"), 64)
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "invalid argument to lat")
	}
	currentLng, err := strconv.ParseFloat(c.QueryParam("lng"), 64)
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "invalid argument to lng")
	}

	googleAPIKey := viper.GetString("GOOGLE_API_KEY")
	lat1, lng1, err := utils.GetCoordinates(address.ZIP, googleAPIKey)
	if err != nil {
		return err
	}

	if currentLat != 0 && currentLng != 0 {
		// checks if current user location is not too far away from address
		distance, err := utils.GetDistance(currentLat, currentLng, lat1, lng1, googleAPIKey)
		if err != nil {
			return err
		}
		if distance > 5000 {
			return echo.NewHTTPError(http.StatusBadRequest, "sorry, the purchase could not be completed: the billing address and regular address are incompatible")
		}
	}

	lat2, lng2, err := utils.GetCoordinates(billingAddress.BillingZIP, googleAPIKey)
	if err != nil {
		return err
	}

	// checks if address and billing address aren't too far away from each other
	distance, err := utils.GetDistance(lat1, lng1, lat2, lng2, googleAPIKey)
	if err != nil {
		return err
	}
	if distance > 5000 {
		return echo.NewHTTPError(http.StatusBadRequest, "sorry, the purchase could not be completed: the billing address and regular address are incompatible")
	}

	var mapResponse map[string]interface{}
	mapResponse = make(map[string]interface{})
	mapResponse["status"] = "ok"
	c.JSON(http.StatusOK, mapResponse)
	return nil
}
