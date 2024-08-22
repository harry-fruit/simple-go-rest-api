package response

import (
	"encoding/json"
	"net/http"
)

type HTTPResponse struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
}

type HTTPResponseInterface interface {
	Send(w http.ResponseWriter)
}

func (h *HTTPResponse) Send(w http.ResponseWriter) {
	w.WriteHeader(h.StatusCode)

	if h.StatusCode == http.StatusNoContent {
		return
	}

	json.NewEncoder(w).Encode(h)
}
