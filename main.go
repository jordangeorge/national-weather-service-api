package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	httpSwagger "github.com/swaggo/http-swagger"

	"api/docs"
	"api/internal"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// API endpoint
	r.Get("/weather-today", internal.GetWeatherTodayHandler)

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "National Weather Service API"
	docs.SwaggerInfo.Description = "This API returns weather forecasts based on latitude and longitude."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"

	// serve Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	port := ":8080"
	fmt.Printf("Server running on %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}
