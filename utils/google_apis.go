package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lualfe/ecommerce/models"
)

// GetCoordinates gets zip codes coordinates from google api
func GetCoordinates(zip, apiKey string) (lat, lng float64, err error) {
	coordinatesURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%v&key=%v", zip, apiKey)
	resp, err := http.Get(coordinatesURL)
	if err != nil {
		return 0, 0, echo.NewHTTPError(resp.StatusCode, "could not get response for provided address: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, 0, echo.NewHTTPError(http.StatusInternalServerError, "could not process response body: ", err)
	}

	var geoCode *models.GeoCode
	err = json.Unmarshal(body, &geoCode)
	if err != nil {
		return 0, 0, echo.NewHTTPError(http.StatusInternalServerError, "could not process json from response: ", err)
	}
	if len(geoCode.Results) > 0 {
		lat = geoCode.Results[0].Geometry.Location.Lat
		lng = geoCode.Results[0].Geometry.Location.Lng
	}
	return
}

// GetDistance get distance between two coordinates
func GetDistance(lat1, lng1, lat2, lng2 float64, apiKey string) (int, error) {
	distanceURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/distancematrix/json?origins=%v,%v&destinations=%v,%v&key=%v", lat1, lng1, lat2, lng2, apiKey)
	resp, err := http.Get(distanceURL)
	if err != nil {
		return 0, echo.NewHTTPError(resp.StatusCode, "could not get response for provided coordinates: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusInternalServerError, "could not process response body: ", err)
	}
	var distanceMatrix *models.DistanceMatrix
	err = json.Unmarshal(body, &distanceMatrix)
	if err != nil {
		return 0, echo.NewHTTPError(http.StatusInternalServerError, "could not process json from response: ", err)
	}
	var distance int
	if len(distanceMatrix.Rows) > 0 {
		if len(distanceMatrix.Rows[0].Elements) > 0 {
			distance = distanceMatrix.Rows[0].Elements[0].Distance.Value
		}
	}
	return distance, nil
}
