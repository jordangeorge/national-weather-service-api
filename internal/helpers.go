package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func defineTemperatures(temp int) string {
	switch {
	case temp <= 55:
		return "cold"
	case temp <= 75:
		return "moderate"
	default:
		return "hot"
	}
}

func validateCoordinates(latitude string, longitude string) error {
	// empty check
	if latitude == "" {
		return fmt.Errorf("latitude missing from the query parameters")
	}
	if longitude == "" {
		return fmt.Errorf("longitude missing from the query parameters")
	}

	// number check
	latNum, err := strconv.ParseFloat(latitude, 64)
	if err != nil {
		return fmt.Errorf("latitude is not a number")
	}
	lonNum, err := strconv.ParseFloat(longitude, 64)
	if err != nil {
		return fmt.Errorf("longitude is not a number")
	}

	// range check
	if latNum < -90 || latNum > 90 {
		return fmt.Errorf("latitude out of range")
	}
	if lonNum < -180 || lonNum > 180 {
		return fmt.Errorf("longitude out of range")
	}

	return nil
}

func writeHTTPErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	response := ErrorResponse{
		Status: statusCode,
		Error:  message,
	}
	json.NewEncoder(w).Encode(response)
}
