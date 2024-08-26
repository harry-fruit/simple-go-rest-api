package handlers

import (
	"encoding/json"
	"net/http"
)

type HTTPResponse struct {
	Data        interface{} `json:"data"`
	StatusCode  int         `json:"status_code"`
	Message     string      `json:"message"`
	contentType string
}

type HTTPResponseInterface interface {
	Send(w http.ResponseWriter)
}

func (h *HTTPResponse) Send(w http.ResponseWriter) {

	if h.contentType != "" {
		w.Header().Set("Content-Type", h.contentType)
	} else {
		w.Header().Set("Content-Type", "application/json")
	}

	w.WriteHeader(h.StatusCode)

	if h.StatusCode == http.StatusNoContent {
		return
	}

	json.NewEncoder(w).Encode(h)
}
