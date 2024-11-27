package shipping

import (
	"database/sql"
	"ecommers/helper"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"go.uber.org/zap"
)

type ShippingRepository struct {
	DB     *sql.DB
	Logger *zap.Logger
}

func NewShippingRepository(Logger *zap.Logger, DB *sql.DB) ShippingRepository {
	return ShippingRepository{
		DB,
		Logger,
	}
}

func (s *ShippingRepository) ShippingCounting(longlat string) (float64, error) {

	baseURL, err := url.Parse("http://router.project-osrm.org/route/v1/driving/")
	if err != nil {
		s.Logger.Error("Error parsing base URL: " + err.Error())
		return 0, err
	}
	originCoordinate := "107.4770867,-6.1980139"
	routeURL := fmt.Sprintf("%s%s;%s?overview=false", baseURL, originCoordinate, longlat)

	header := make(http.Header)
	data, err := helper.HTTPRequest("GET", header, routeURL, nil)
	if err != nil {
		s.Logger.Error("Error making HTTP request: " + err.Error())
		return 0, errors.New("error http request direction")
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		s.Logger.Error("Error parsing OSRM response: " + err.Error())
		return 0, err
	}

	routes, ok := dataMap["routes"].([]interface{})
	if !ok || len(routes) == 0 {
		s.Logger.Error("No routes found in OSRM response")
		return 0, errors.New("no routes found in OSRM response")
	}

	route, ok := routes[0].(map[string]interface{})
	if !ok {
		s.Logger.Error("Error decoding route data")
		return 0, errors.New("error decoding route data")
	}

	distance, ok := route["distance"].(float64)
	if !ok {
		s.Logger.Error("Error decoding distance")
		return 0, errors.New("error decoding distance")
	}

	return distance, nil
}
