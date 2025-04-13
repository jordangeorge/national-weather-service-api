package internal

import (
	"encoding/json"
	"net/http"
)

// GetWeatherTodayHandler godoc
// @Summary      Get today's weather
// @Description  Returns short forecast and temperature description
// @Tags         weather-today
// @Accept       json
// @Produce      json
// @Param        latitude   query     string  true  "Latitude"
// @Param        longitude  query     string  true  "Longitude"
// @Success      200  {array}   SuccessResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /weather-today [get]
func GetWeatherTodayHandler(w http.ResponseWriter, r *http.Request) {
	// get latitude and longitude from the request
	latitude := r.URL.Query().Get("latitude")
	longitude := r.URL.Query().Get("longitude")

	// validate the latitude and longitude
	err := validateCoordinates(latitude, longitude)
	if err != nil {
		writeHTTPErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	// get the forecasts
	forecasts, err := getWeatherToday(latitude, longitude)
	if err != nil {
		writeHTTPErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	// send success response
	if len(forecasts) > 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(forecasts)
		return
	}

	// send error response
	writeHTTPErrorResponse(w, http.StatusNotFound, "No forecasts found for today")
}
