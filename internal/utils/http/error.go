package handlers

import (
	"net/http"
	"strings"
)

const LoginInUse = "LOGIN_IN_USE"
const BadPayload = "BAD_PAYLOAD"

type HTTPError struct {
	StatusCode int
	Message    string
	ErrorType  string
}

func (he HTTPError) SendError(w http.ResponseWriter) {

	httpResponse := &HTTPResponse{
		StatusCode: he.getStatusCode(),
		Message:    he.getMessage(),
		Data:       nil,
	}

	//Persist in logs
	httpResponse.Send(w)
}

func (he HTTPError) getMessage() string {

	if he.Message != "" {
		return he.Message
	}

	switch he.ErrorType {
	case LoginInUse:
		return "login already in use"
	case BadPayload:
		return "bad payload"
	default:
		return strings.ToLower(http.StatusText(he.getStatusCode()))
	}

}

func (he HTTPError) getStatusCode() int {
	if he.StatusCode != 0 {
		return he.StatusCode
	}
	return http.StatusInternalServerError
}
