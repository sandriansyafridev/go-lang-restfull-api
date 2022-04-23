package response

import "strings"

type WebResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

type EmptyObject struct {
}

func BuildResponseSuccess(message string, data interface{}) WebResponse {
	return WebResponse{
		Status:  true,
		Message: message,
		Data:    data,
		Errors:  nil,
	}
}

func BuildResponseError(message string, errors string, data interface{}) WebResponse {

	splitError := strings.Split(errors, "\n")

	return WebResponse{
		Status:  false,
		Message: message,
		Data:    data,
		Errors:  splitError,
	}
}
