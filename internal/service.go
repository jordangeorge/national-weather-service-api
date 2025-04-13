package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func getWeatherToday(latitude string, longitude string) ([]SuccessResponse, error) {
	// get points data from api response
	pointsURL := fmt.Sprintf(
		"https://api.weather.gov/points/%s,%s",
		latitude,
		longitude,
	)
	pointsResp, err := http.Get(pointsURL)
	if err != nil || pointsResp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch points data: %s", err.Error())
	}
	defer pointsResp.Body.Close()

	// decode to points model
	pointsData := PointsData{}
	if err := json.NewDecoder(pointsResp.Body).Decode(&pointsData); err != nil {
		return nil, fmt.Errorf("failed to decode points data: %s", err.Error())
	}

	// get forecast data from api response
	forecastURL := pointsData.Properties.ForecastURL
	forecastResp, err := http.Get(forecastURL)
	if err != nil || forecastResp.StatusCode != 200 {
		return nil, fmt.Errorf("failed to fetch forecast data: %s", err.Error())
	}
	defer forecastResp.Body.Close()

	// decode to forecast model
	forecastData := ForecastData{}
	if err := json.NewDecoder(forecastResp.Body).Decode(&forecastData); err != nil {
		return nil, fmt.Errorf("failed to decode forecast data: %s", err.Error())
	}

	// save today's date
	today := time.Now().Format("2006-01-02")

	// get short forecast for today from forecast model
	forecasts := []SuccessResponse{}
	for _, period := range forecastData.Properties.Periods {
		st := period.StartTime[:10]
		if st == today {
			forecast := SuccessResponse{
				Time:        period.Name,
				Forecast:    period.ShortForecast,
				Description: defineTemperatures(period.Temperature),
			}
			forecasts = append(forecasts, forecast)
		}
	}

	return forecasts, nil
}
