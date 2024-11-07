package api

import (
	"context"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"time"
)


// APIHandler is a type to give the api functions below access to a common logger
// any any other shared objects
type APIHandler struct {
	// Zerolog was chosen as the default logger, but you can replace it with any logger of your choice
	logger zerolog.Logger

	// Note: if you need to pass in a client for your database, this would be a good place to include it
}

func NewAPIHandler() *APIHandler {
	output := zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	return &APIHandler{logger: logger}
}

func (h *APIHandler) WithLogger(logger zerolog.Logger) *APIHandler {
	h.logger = logger
	return h
}

// Fetches the latest sports news articles.
// Retrieve sports news
func (h *APIHandler) GetSportsNews(ctx context.Context) (Response, error) {
	// TODO: implement the GetSportsNews function to return the following responses

	// return NewResponse(200, SportsNewsResponse{}, "application/json", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"getSportsNews operation has not been implemented yet"}, "application/json", nil), nil
}

// Fetches the latest weather updates.
// Retrieve weather updates
func (h *APIHandler) GetWeatherUpdates(ctx context.Context) (Response, error) {
	// TODO: implement the GetWeatherUpdates function to return the following responses

	// return NewResponse(200, WeatherResponse{}, "application/json", responseHeaders), nil

	// return NewResponse(404, {}, "", responseHeaders), nil

	// return NewResponse(500, {}, "", responseHeaders), nil

	return NewResponse(http.StatusNotImplemented, ErrorMsg{"getWeatherUpdates operation has not been implemented yet"}, "application/json", nil), nil
}

