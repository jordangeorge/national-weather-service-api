package internal

// PointsData : starting point of parsing data from the /points/{latitude},{longitude} endpoint response
type PointsData struct {
	Properties PointsProperties `json:"properties"`
}

type PointsProperties struct {
	ForecastURL string `json:"forecast"`
}

// ForecastData : starting point of parsing the data from the /gridpoints/{office}/{latitude},{longitude}/forecast endpoint response
type ForecastData struct {
	Properties GridpointsForecastProperty `json:"properties"`
}

type GridpointsForecastProperty struct {
	Periods []GridpointForecastPeriod `json:"periods"`
}

type GridpointForecastPeriod struct {
	Name          string `json:"name"`
	Temperature   int    `json:"temperature"`
	ShortForecast string `json:"shortForecast"`
	StartTime     string `json:"startTime"`
}

// SuccessResponse is this API's success response
type SuccessResponse struct {
	Time        string `json:"time"`
	Forecast    string `json:"forecast"`
	Description string `json:"description"`
}

// ErrorResponse is this API's error response
type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}
