package handlers

import (
	"net/http"
)

const LoginInUse = "LOGIN_IN_USE"

func SendError(w http.ResponseWriter, err error) {
	statusCode, message := getStatusCodeAndMessage(err)

	httpResponse := &HTTPResponse{
		StatusCode: statusCode,
		Message:    message,
		Data:       nil,
	}

	httpResponse.Send(w)
}

func getStatusCodeAndMessage(err error) (int, string) {
	var statusCode int
	var message string

	switch err.Error() {
	case LoginInUse:
		statusCode = http.StatusConflict
		message = "login in use"
	default:
		statusCode = http.StatusInternalServerError
		message = "internal server error"
	}

	return statusCode, message
}
