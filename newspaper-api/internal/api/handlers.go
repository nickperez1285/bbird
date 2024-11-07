package api

import (
	"net/http"
)

// HandleGetSportsNews handles parsing input to pass to the GetSportsNews operation and sends responses back to the client
func (h *APIHandler) HandleGetSportsNews(w http.ResponseWriter, r *http.Request) {
	var err error
	response, err := h.GetSportsNews(r.Context())
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetSportsNews returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetSportsNews was unable to send it's response, err: %s", err)
	}
}

// HandleGetWeatherUpdates handles parsing input to pass to the GetWeatherUpdates operation and sends responses back to the client
func (h *APIHandler) HandleGetWeatherUpdates(w http.ResponseWriter, r *http.Request) {
	var err error
	response, err := h.GetWeatherUpdates(r.Context())
	if err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetWeatherUpdates returned err: %s", err)
	}

	if err = response.Send(w); err != nil {
		ErrorResponse(http.StatusInternalServerError, w)
		h.logger.Error().Msgf("GetWeatherUpdates was unable to send it's response, err: %s", err)
	}
}

